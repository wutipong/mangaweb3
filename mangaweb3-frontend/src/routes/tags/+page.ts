import type { PageLoad } from './$types';
import { variables } from '$lib/variables';
import type { Tag } from '$lib/tag';

interface listRequest {
    favorite_only: boolean
    page: number,
    item_per_page: number,
}

interface listResponse {
    tags: Tag[]
}

export const load: PageLoad = async ({ fetch, url }) => {
    const tagListURL = new URL("/tag/list", variables.basePath);
    const request: listRequest = {
        favorite_only: false,
        page: 0,
        item_per_page: 30
    }

    if (url.searchParams.has('favorite')) {
        request.favorite_only = url.searchParams.get('favorite') == "true"
    }
    if (url.searchParams.has('page')) {
        const v = url.searchParams.get('page')
        if (v != null)
            request.page = parseInt(v)
    }
    if (url.searchParams.has('item_per_page')) {
        const v = url.searchParams.get('item_per_page')
        if (v != null)
            request.item_per_page = parseInt(v)
    }

    const response = await fetch(tagListURL, { method: 'POST', body: JSON.stringify(request) });

    const obj = await response.json() as listResponse;

    return {
        page: request.page,
        tags: obj?.tags as Tag[],
        total_page: obj.tags
    };
};