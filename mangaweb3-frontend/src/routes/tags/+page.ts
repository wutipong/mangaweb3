import type { PageLoad } from './$types';
import { variables } from '$lib/variables';

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

    const tagListURL = new URL("/tag/list", variables.basePath);
    const response = await fetch(tagListURL, { method: 'POST' });

    const obj = await response.json() as listResponse;

    return {
        tags: obj?.tags as Tag[]
    };
};