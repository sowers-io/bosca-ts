import { SessionInterceptor, Subject, SubjectFinder } from './subject_finder';
export declare class HttpSubjectFinder implements SubjectFinder {
    private readonly endpoint;
    private readonly serviceAccountId;
    private readonly serviceAccountTokenHeader;
    private readonly interceptor;
    constructor(endpoint: string, serviceAccountId: string, serviceAccountToken: string, interceptor: SessionInterceptor);
    findSubject(fromCookie: boolean, authorization: string): Promise<Subject>;
}
export declare class HttpSessionInterceptor implements SessionInterceptor {
    getSubjectId(response: Response): Promise<string | null>;
}
//# sourceMappingURL=http_subject_finder.d.ts.map