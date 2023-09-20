import { useState } from "react";
import "./App.css";
import { useClient } from "./hooks/useClient";

import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// Import service definition that you want to connect to.
import { GreetService } from "../gen/greet_connect";
import { createAsyncIterable } from "@connectrpc/connect/protocol";
import { PartialMessage } from "@bufbuild/protobuf";
import { GreetRequest } from "../gen/greet_pb";

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
const transport = createConnectTransport({
  // baseUrl: "https://demo.connectrpc.com",
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(GreetService, transport);

function App() {
  const [inputValue, setInputValue] = useState("");
  const [messages, setMessages] = useState<
    {
      fromMe: boolean;
      message: string;
    }[]
  >([]);
  const client = useClient(GreetService);
  return (
    <>
      <ol>
        {messages.map((msg, index) => (
          <li key={index}>
            {`${msg.fromMe ? "ME:" : "ELIZA:"} ${msg.message}`}
          </li>
        ))}
      </ol>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          // Clear inputValue since the user has submitted.
          setInputValue("");
          // Store the inputValue in the chain of messages and
          // mark this message as coming from "me"
          setMessages((prev) => [
            ...prev,
            {
              fromMe: true,
              message: inputValue,
            },
          ]);
          const a = new GreetRequest({ name: "test" });
          const b = new GreetRequest({ name: "test2" });
          const d = [a, b];
          const test = createAsyncIterable<PartialMessage<GreetRequest>>(d);
          const response = await client.greet(test);
          console.log(response);
          // setMessages((prev) => [
          //   ...prev,
          //   {
          //     fromMe: false,
          //     message: response,
          //   },
          // ]);
        }}
      >
        <input
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
        />
        <button type="submit">Send</button>
      </form>
    </>
  );
}

export default App;
