import * as messages from './guruguru_pb'
import * as services from './guruguru_grpc_pb'
import * as grpc from 'grpc'
import { promisify } from 'util'

const getConfig = () => {
  return {
    name: process.env.MEMBER_NAME || 'node',
    hostPort: process.env.HOST_PORT || '0.0.0.0:5000',
    publicHostPort: process.env.PUBLIC_HOST_PORT || 'localhost:5000',
    bossHostPort: process.env.BOSS_HOST_PORT || 'boss:5000'
  }
}

const join = async config => {
  const client = new services.BossServiceClient(
    config.bossHostPort,
    grpc.credentials.createInsecure()
  )
  const req = new messages.JoinRequest()
  const joiningMember = new messages.Member()
  joiningMember.setName(config.name)
  joiningMember.setHostPort(config.publicHostPort)
  req.setJoiningMember(joiningMember)

  const res = await promisify(client.join.bind(client))(req)

  if (!res.getOk()) {
    throw Error('colud not join')
  }
}

const poke = async (to, config) => {
  const client = new services.MemberServiceClient(
    to.hostPort,
    grpc.credentials.createInsecure()
  )
  const req = new messages.PokeRequest()
  const fromMember = new messages.Member()
  fromMember.setName(config.name)
  fromMember.setHostPort(config.publicHostPort)
  req.setFromMember(fromMember)
  req.setMessage('Give me callbacks.')

  const res = await promisify(client.poke.bind(client))(req)

  if (!res.getOk()) {
    throw Error('colud not poke')
  }
}

const handleSetNext = state => {
  return (call, callback) => {
    const next = call.request.getNextMember()
    state.next = {
      name: next.getName(),
      hostPort: next.getHostPort()
    }

    console.log(`Set next to ${state.next.name}(${state.next.hostPort})`)

    const res = new messages.SetNextResponse()
    res.setOk(true)
    callback(null, res)
  }
}

const handlePoke = state => {
  return (call, callback) => {
    const fromMember = call.request.getFromMember()

    console.log(
      `Got message "${call.request.getMessage()}" from ${fromMember.getName()}. Hey ${
        state.next.name
      }!`
    )

    poke(state.next, state.config)

    const res = new messages.PokeResponse()
    res.setOk(true)
    callback(null, res)
  }
}

const run = async () => {
  const config = getConfig()
  console.log(`Starting... I'm ${config.name}`)

  await join(config)

  const state = { config: config }
  const server = new grpc.Server()
  server.addService(services.MemberServiceService, {
    setNext: handleSetNext(state),
    poke: handlePoke(state)
  })
  server.bind(config.hostPort, grpc.ServerCredentials.createInsecure())
  server.start()
}

run()
