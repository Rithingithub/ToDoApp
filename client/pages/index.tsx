import {
  Box,
  Button,
  Flex,
  Heading,
  Input,
  InputGroup,
  InputRightAddon,
  Stack,
} from "@chakra-ui/react";
import React, { useState } from "react";

const Index = () => {
  const [inputValue, setInputValue] = useState<string>("");
  const [data, setData] = useState<string[]>([]);
  const handleInputEvent = (e: any) => {
    setInputValue(e.target?.value);
    console.log(inputValue);
  };
  const getData = () => {
    setData([...data, inputValue]);
    console.log(data);
  };
  return (
    <Flex
      flexDirection="column"
      width="100wh"
      height="100vh"
      alignItems="center"
      backgroundColor="blue.200"
    >
      <Stack
        flexDir="column"
        mb="4"
        justifyContent="center"
        alignItems="center"
      >
        <Box backgroundColor="white" mt="5" padding="5" width="100vh">
          <Heading mt="5" mb="10" textAlign="center">
            TODO
          </Heading>
          <InputGroup>
            <Input
              placeholder="title"
              onChange={handleInputEvent}
              value={inputValue}
            />
            {/* <InputRightAddon children="+234" /> */}
            <Button onClick={getData}> ADD</Button>
          </InputGroup>
        </Box>
      </Stack>
    </Flex>
  );
};

export default Index;
function setInputValue(arg0: HTMLInputElement) {
  throw new Error("Function not implemented.");
}
