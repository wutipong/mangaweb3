import type { PageLoad } from './$types';
import { getBackendBaseURL } from '$lib/config'

interface Request {
    path: string
}

interface Response {
    browse_url: string
    favorite: boolean
    indices: number[]
    tags: string[]
}

export const load: PageLoad = async ({ fetch, params, url}) => {
    params
    const name = url.searchParams.get("name");
    const request = {
        path: name,
    }

    const tagListURL = new URL("/view", getBackendBaseURL());
    const response = await fetch(tagListURL, { method: 'POST', body: JSON.stringify(request) });

    const obj = await response.json() as Response;

    return {
        name: name !== null? name: "",
        browse_url: obj.browse_url,
        favorite: obj.favorite,
        indices: obj.indices,
        tags: obj.tags,
    };
};