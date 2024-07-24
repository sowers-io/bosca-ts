// Original is from p-limit, modified to not have dependency and to not require changes to module configuration
export default function pLimit(concurrency: number) {
  let queue = []
  let activeCount = 0

  const resumeNext = () => {
    if (activeCount < concurrency && queue.length > 0) {
      queue.shift()()
      // Since `pendingCount` has been decreased by one, increase `activeCount` by one.
      activeCount++
    }
  }

  const next = () => {
    activeCount--

    resumeNext()
  }

  const run = async (function_, resolve, arguments_) => {
    const result = (async () => function_(...arguments_))()

    resolve(result)

    try {
      await result
    } catch {}

    next()
  }

  const enqueue = (function_, resolve, arguments_) => {
    // Queue `internalResolve` instead of the `run` function
    // to preserve asynchronous context.
    new Promise((internalResolve) => {
      queue.push(internalResolve)
    }).then(run.bind(undefined, function_, resolve, arguments_))

    ;(async () => {
      // This function needs to wait until the next microtask before comparing
      // `activeCount` to `concurrency`, because `activeCount` is updated asynchronously
      // after the `internalResolve` function is dequeued and called. The comparison in the if-statement
      // needs to happen asynchronously as well to get an up-to-date value for `activeCount`.
      await Promise.resolve()

      if (activeCount < concurrency) {
        resumeNext()
      }
    })()
  }

  const generator = (function_, ...arguments_) =>
    new Promise((resolve) => {
      enqueue(function_, resolve, arguments_)
    })

  Object.defineProperties(generator, {
    activeCount: {
      get: () => activeCount,
    },
    pendingCount: {
      get: () => queue.length,
    },
    clearQueue: {
      value() {
        queue = []
      },
    },
    concurrency: {
      get: () => concurrency,

      set(newConcurrency) {
        concurrency = newConcurrency

        queueMicrotask(() => {
          // eslint-disable-next-line no-unmodified-loop-condition
          while (activeCount < concurrency && queue.length > 0) {
            resumeNext()
          }
        })
      },
    },
  })

  return generator
}
