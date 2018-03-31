$:.unshift File.dirname(__FILE__)
require 'grpc'
require 'guruguru_services_pb'

STDOUT.sync = true


Mem = Struct.new(:name, :host_port, keyword_init: true) # Member is already defined
Config = Struct.new(:name, :host_port, :public_host_port, :boss_host_port, keyword_init: true)

def _config
  Config.new(
    name: ENV.fetch('MEMBER_NAME', 'ruby'),
    host_port: ENV.fetch('HOST_PORT', '0.0.0.0:5000'),
    public_host_port: ENV.fetch('PUBLIC_HOST_PORT', 'ruby:5000'),
    boss_host_port: ENV.fetch('BOSS_HOST_PORT', 'boss:5000')
  )
end

def callJoin(config)
  stub = BossService::Stub.new(config.boss_host_port, :this_channel_is_insecure)
  res = stub.join(
    JoinRequest.new(
      joining_member: Member.new(name: config.name, host_port: config.public_host_port)
    )
  )
  raise 'could not join' unless res.ok
end

def callPoke(to, config)
  stub = MemberService::Stub.new(to.host_port, :this_channel_is_insecure)
  res = stub.poke(
    PokeRequest.new(
      from_member: Member.new(name: config.name, host_port: config.public_host_port),
      message: 'Naming is important'
    )
  )
  raise 'could not poke' unless res.ok
end

class Server < MemberService::Service
  def initialize(config)
    @config = config
    @mutex = Mutex.new
    @next = nil
  end

  def set_next(req, _call)
    next_member = req.next_member

    @mutex.synchronize {
      @next = Mem.new(name: next_member.name, host_port: next_member.host_port)
    }

    SetNextResponse.new(ok: true)
  end

  def poke(req, _call)
    from_member = req.from_member
    message = req.message
    puts("Got message \"#{message}\" from #{from_member.name}. Hey #{@next.name}!")

    callPoke(@next, @config) # not async
    PokeResponse.new(ok: true)
  end
end

def run
  c = _config
  puts("Starting I'm ... #{c.name}")
  sleep(1) # wait for boss
  callJoin(c)

  s = GRPC::RpcServer.new
  s.add_http2_port(c.host_port, :this_port_is_insecure)
  s.handle(Server.new(c))
  s.run_till_terminated
end

run
