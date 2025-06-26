import { render } from 'ink'
import { App } from './app'

async function main() {
  const instance = render(null)
  // enter alternate buffer
  await write('\x1b[?1049h')
  instance.rerender(<App />)
  await instance.waitUntilExit()
  // exit alternate buffer
  await write('\x1b[?1049l')
  // close the app
  process.exit(0)
}
main()

async function write(content: string) {
  return new Promise<void>((resolve, reject) => {
    process.stdout.write(content, (error) => {
      if (error) {
        reject(error)
      } else {
        resolve()
      }
    })
  })
}
