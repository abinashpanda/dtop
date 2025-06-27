import { z } from 'zod'
import { unix } from '../../lib/socket'

export async function getContainers() {
  const res = await fetch('http://localhost/containers/json?all=true', {
    unix,
  })
  const data = z
    .object({
      Id: z.string(),
      Names: z.string().array(),
      Image: z.string(),
      ImageID: z.string(),
      Command: z.string(),
      Ports: z
        .object({
          IP: z.string().array(),
          PrivatePort: z.number(),
          PublicPort: z.number(),
        })
        .array(),
      State: z.enum([
        'created',
        'running',
        'paused',
        'restarting',
        'exited',
        'removing',
        'dead',
      ]),
      Labels: z.object({ name: z.string().optional() }),
    })
    .array()
    .parse(await res.json())
  return data
}
