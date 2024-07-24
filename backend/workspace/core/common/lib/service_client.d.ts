import { PromiseClient } from '@connectrpc/connect';
import { ServiceType } from '@bufbuild/protobuf';
export declare function useServiceAccountClient<T extends ServiceType>(service: T): PromiseClient<T>;
export declare function useClient<T extends ServiceType>(service: T): PromiseClient<T>;
//# sourceMappingURL=service_client.d.ts.map