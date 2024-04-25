export async function handlePromise<T>(
    promise: Promise<T>
): Promise<[T, Error]> {
    if (promise && typeof promise.then === 'function') {
        return promise
            .then((data: T) => [data, undefined])
            .catch((err: Error) => [undefined, err]);
    }
    throw new Error('only promise is allowed');
}
