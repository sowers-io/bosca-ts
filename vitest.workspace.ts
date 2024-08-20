import { defineWorkspace } from 'vitest/config'

export default defineWorkspace([
  "./workspace/core/workflow-activities-api/vitest.config.ts",
  "./workspace/bible/processor/vitest.config.ts",
  "./workspace/workflow/ai/vitest.config.ts"
])
