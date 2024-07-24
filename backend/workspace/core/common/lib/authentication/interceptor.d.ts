import type { Interceptor } from '@connectrpc/connect';
import { SubjectFinder, Subject } from './subject_finder';
export declare const SubjectKey: import("@connectrpc/connect").ContextKey<Subject>;
export declare function newAuthenticationInterceptor(subjectFinder: SubjectFinder): Interceptor;
//# sourceMappingURL=interceptor.d.ts.map