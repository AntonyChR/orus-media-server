import { useState } from 'react';

/**
 * Custom hook that wraps a fetcher function and returns the data, loading state and error.
 * @param fetcher - Function that fetches the data.
 * @returns data - The fetched data.
 * @returns loading - Whether the fetch is in progress.
 * @returns error - The error that occurred during the fetch.
 * @returns makeRequest - Function that triggers the fetch.
 */
export function useWrapFetch<T>(fetcher: () => Promise<T>):{
    data: T | null;
    loading: boolean;
    error: Error | null;
    makeRequest: ()=>void;
} {
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
