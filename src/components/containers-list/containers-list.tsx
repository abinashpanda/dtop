import { Box, Text } from 'ink'

type ContainersListProps = Omit<React.ComponentProps<typeof Box>, 'borderStyle'>

export default function ContainersList({ ...rest }: ContainersListProps) {
  return (
    <Box {...rest} borderStyle="round">
      <Box position="absolute" marginTop={-1} marginLeft={1}>
        <Text>Containers</Text>
      </Box>
    </Box>
  )
}
