import { useState } from "react";
import {
  Flex,
  Text,
  Heading,
  Input,
  Button,
  InputGroup,
  Stack,
  chakra,
  Box,
  Link,
  Avatar,
  FormControl,
  FormHelperText,
  InputRightElement,
} from "@chakra-ui/react";

const App = () => {
  const [showPassword, setShowPassword] = useState(false);

  const handleShowClick = () => setShowPassword(!showPassword);
  const [issignUp, setsignUp] = useState<boolean>(false);

  return (
    <Flex
      flexDirection="column"
      width="100wh"
      height="100vh"
      backgroundColor="whiteAlpha.900"
      justifyContent="center"
      alignItems="center"
    >
      <Stack
        flexDir="column"
        mb="2"
        justifyContent="center"
        alignItems="center"
      >
        <Box minW={{ base: "90%", md: "468px" }}>
          <form>
            <Stack spacing={4} p="1rem" backgroundColor="whiteAlpha.900">
              <Heading mb="4">Sign in</Heading>
              <FormControl>
                <Text>Username</Text>
                <InputGroup>
                  <Input variant="flushed" />
                </InputGroup>
              </FormControl>
              <FormControl>
                <Text>Password</Text>
                <InputGroup>
                  <Input
                    variant="flushed"
                    type={showPassword ? "text" : "password"}
                  />
                  <InputRightElement width="4.5rem">
                    <Button h="1.75rem" size="sm" onClick={handleShowClick}>
                      {showPassword ? "Hide" : "Show"}
                    </Button>
                  </InputRightElement>
                </InputGroup>
                <FormHelperText textAlign="left" color="blue">
                  <Link>forgot password?</Link>
                </FormHelperText>
              </FormControl>
              <Button
                mt="4"
                type="submit"
                variant="outline"
                colorScheme="black"
              >
                Log in
              </Button>
              <Button
                type="submit"
                variant="solid"
                border="2px"
                borderColor="black"
                colorScheme="teal"
              >
                Create an account
              </Button>
            </Stack>
          </form>
        </Box>
      </Stack>
    </Flex>
  );
};

export default App;
