import Content from "@/components/content";
import { Suspense } from "react";

export default function Page() {
  return (
    <Suspense>
      <Content />
    </Suspense>
  );
}
