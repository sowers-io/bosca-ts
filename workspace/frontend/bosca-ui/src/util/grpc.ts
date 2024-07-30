"use server";

import "server-only";

import { getAuthorizationToken } from "@/util/security";
import { Interceptor } from "@connectrpc/connect";
import { setClientInterceptor } from "@bosca/common";

export async function initializeClient() {
  const authorization: Interceptor = (next) => async (req) => {
    const token = getAuthorizationToken();
    if (token) {
      req.header.set("Cookie", "ory_kratos_session=" + token);
    }
    return await next(req);
  };
  setClientInterceptor(authorization);
}
