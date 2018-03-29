import * as messages from './guruguru_pb'
import * as services from './guruguru_grpc_pb'

const getConfig = () => {
  return {
    name: process.env.MEMBER_NAME || 'node',
    host_port: process.env.HOST_PORT || '0.0.0.0:5000',
    public_host_port: process.env.PUBLIC_HOST_PORT || 'localhost:5000',
    boss_host_port: process.env.BOSS_HOST_PORT || 'boss:5000'
  }
}

const join = async config => {
  const client = new services.BossServiceClient(config.boss_host_port)
  const req = new messages.JoinRequest()
  const joiningMember = new messages.Message()
  joiningMember.setName(config.name)
  joiningMember.setHostPort(config.public_host_port)
  req.setJoiningMember(joiningMember)

  const res = await client.join(req)
  console.log(res)

  if (!res.ok) {
    throw Error('colud not join')
  }
}

const run = () => {
  const config = getConfig()

  await join(config)
}

run()
