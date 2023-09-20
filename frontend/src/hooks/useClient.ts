import { useMemo } from "react";
import { ServiceType } from "@bufbuild/protobuf";
import {
  createConnectTransport,
  createGrpcWebTransport,
} from "@connectrpc/connect-web";
import { createPromiseClient, PromiseClient } from "@connectrpc/connect";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  // generate stream http client
});
// createGrpcWebTransport({

// })

export function useClient<T extends ServiceType>(service: T): PromiseClient<T> {
  return useMemo(() => createPromiseClient(service, transport), [service]);
}
