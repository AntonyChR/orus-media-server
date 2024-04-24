import { useState } from 'react';

/**
 * Custom hook that wraps a fetcher function and handles the state of the fetch operation.
 * @template T The type of data returned by the fetcher function.
 * @param {() => Promise<T>} fetcher The fetcher function that returns a Promise of type T.
 * @returns {Object} An object containing the data, loading state, error, and a function to make the request.
 */
export function useWrapFetch<T>(fetcher: () => Promise<T>) {
    const [data, setData] = useState<T | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<Error | null>(null);

    /**
     * Makes the request using the fetcher function and updates the state accordingly.
     */
    const makeRequest =  () => {
        setLoading(true);
        fetcher()
            .then((d) => {
                setData(d);
                setLoading(false);
            })
            .catch((e) => {
                setError(e);
                setLoading(false);
            });
    }

    return { data, loading, error, makeRequest };
}
