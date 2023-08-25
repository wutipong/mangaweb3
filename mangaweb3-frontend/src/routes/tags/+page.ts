import type { PageLoad } from './$types';
import {getBackendBaseURL} from '$lib/config'

interface listRequest {
    favorite_only: boolean
}

interface Tag {
    favorite: boolean
    id: number
    name: string
}

interface listResponse {
    tags: Tag[]
}

export const load: PageLoad = async ({ fetch, params }) => {

    const tagListURL = new URL("/tag/list", getBackendBaseURL());
    const response = await fetch(tagListURL, { method: 'POST' });

    const obj = await response.json() as listResponse;

    return {
        tags: obj?.tags as Tag[]
    };
};