import { useState } from 'react';

/**
 * Custom hook that wraps a fetcher function and handles the state for making API requests.
 *
 * @template T - The type of data returned by the fetcher function.
 * @template P - The type of arguments accepted by the fetcher function.
 * @param {Function} fetcher - The fetcher function that makes the API request.
 * @returns {Object} - An object containing the data, loading state, error, and a function to make the request.
 */
// eslint-disable-next-line
export function useWrapFetch<T, P = any>(
    fetcher: (args?: P) => Promise<T>
): {
    data: T | null;
    loading: boolean;
    error: Error | null;
    makeRequest: (args?: P) => void;
} {
    const [data, setData] = useState<T | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<Error | null>(null);

    /**
     * Function to make the API request.
     *
     * @param {P} args - The arguments to be passed to the fetcher function.
     */
    const makeRequest = (args?: P) => {
        setLoading(true);
        fetcher(args)
            .then((d) => {
                setData(d);
                setLoading(false);
            })
            .catch((e) => {
                setError(e);
                setLoading(false);
            });
    };

    return { data, loading, error, makeRequest };
}
