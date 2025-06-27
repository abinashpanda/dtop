import { useQuery } from '@tanstack/react-query'
import { Box, Text } from 'ink'
import { getContainers } from './queries'
import { match } from 'ts-pattern'

type ContainersListProps = Omit<React.ComponentProps<typeof Box>, 'borderStyle'>

export default function ContainersList({ ...rest }: ContainersListProps) {
  const containersListQuery = useQuery({
    queryKey: ['containers-list'],
    queryFn: getContainers,
  })

  return (
    <Box {...rest} borderStyle="round">
      <Box position="absolute" marginTop={-1} marginLeft={1}>
        <Text>Containers</Text>
      </Box>
      <Box>
        {match(containersListQuery)
          .returnType<React.ReactNode>()
          .with({ status: 'error' }, ({ error }) => {
            return <Text>Error: {error.message}</Text>
          })
          .with({ status: 'success' }, ({ data: containers }) => {
            return (
              <Box flexDirection="column" width="100%">
                {containers.slice(0, 5).map((container, index) => (
                  <Box key={container.Id} flexDirection="row">
                    <Text>{container.Labels.name}</Text>
                  </Box>
                ))}
              </Box>
            )
          })
          .otherwise(() => null)}
      </Box>
    </Box>
  )
}
