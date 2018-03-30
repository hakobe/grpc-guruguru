from collections import namedtuple
import os
from concurrent import futures
import time
import grpc
import guruguru_pb2
import guruguru_pb2_grpc

Config = namedtuple('Config', ['name', 'host_port', 'public_host_port', 'boss_host_port'])
Member = namedtuple('Member', ['name', 'host_port'])


def get_config():
    return Config(
        os.getenv('MEMBER_NAME', 'python'),
        os.getenv('HOST_PORT', '0.0.0.0:5000'),
        os.getenv('PUBIC_HOST_PORT', 'python:5000'),
        os.getenv('BOSS_HOST_PORT', 'boss:5000'),
    )


def log(*args):
    print(*args, flush=True)


def join(config):
    channel = grpc.insecure_channel(config.boss_host_port)
    stub = guruguru_pb2_grpc.BossServiceStub(channel)
    res = stub.Join(
        guruguru_pb2.JoinRequest(joining_member=guruguru_pb2.Member(name='python', host_port='python:5000')))
    if not res.ok:
        raise RuntimeError('could not join')


def poke(member, config):
    channel = grpc.insecure_channel(member.host_port)
    stub = guruguru_pb2_grpc.MemberServiceStub(channel)
    # not async
    res = stub.Poke(
        guruguru_pb2.PokeRequest(
            from_member=guruguru_pb2.Member(name=config.name, host_port=config.host_port), message='spam! spam! spam!'))
    if not res.ok:
        raise RuntimeError('could not poke')


class Server(guruguru_pb2_grpc.MemberServiceServicer):
    def __init__(self, config):
        super()
        self.config = config
        self.next = None

    def Poke(self, request, context):
        from_member = request.from_member
        message = request.message
        log(f'Got message "{message}" from {from_member.name}. Hey {self.next.name}!')

        poke(self.next, self.config)
        return guruguru_pb2.PokeResponse(ok=True)

    def SetNext(self, request, context):
        next_member = request.next_member
        self.next = Member(next_member.name, next_member.host_port)  # GIL
        return guruguru_pb2.SetNextResponse(ok=True)


def serve(config):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    guruguru_pb2_grpc.add_MemberServiceServicer_to_server(Server(config), server)
    server.add_insecure_port(config.host_port)
    server.start()
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)


def run():
    config = get_config()
    log(f"Starting... I'm {config.name}")
    time.sleep(1)  # wait boss

    join(config)
    serve(config)


if __name__ == '__main__':
    run()
