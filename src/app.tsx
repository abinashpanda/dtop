import { Box, useApp, useInput } from 'ink'
import { useScreenSize } from './hooks/use-screen-size'
import ContainersList from './components/containers-list'

export function App() {
  const app = useApp()
  useInput(function closeAppOnUserInput(input, key) {
    if (key.escape || (key.ctrl && input === 'c') || input === 'q') {
      app.exit()
    }
  })

  const { width, height } = useScreenSize()

  return (
    <Box
      width={width}
      // height - 1 can the last line would always be rendered when using useApp
      height={height - 1}
      flexDirection="row"
      alignItems="flex-start"
    >
      <ContainersList width="30%" />
    </Box>
  )
}
