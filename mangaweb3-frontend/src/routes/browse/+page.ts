import type { PageLoad } from './$types';
import { getBackendBaseURL } from '$lib/config'

interface browseRequest {
    favorite_only: boolean
    item_per_page: number
    order: ["ascending" | "descending"]
    page: number
    search: string
    sort: ["name" | "createTime"]
    tag: string
}

interface Item {
    create_time: string
    favorite: boolean
    id: number
    is_read: boolean
    name: string
}

interface Page {
    content: string
    is_active: boolean
    is_enabled: boolean
    is_hidden_on_small: boolean
    link_url: string
}

interface browseResponse {
    items: Item[]
    pages: Page[]
    tag_favorite: boolean
    total_page: number
}


export const load: PageLoad = async ({ fetch, params }) => {

    const browseURL = new URL("/browse", getBackendBaseURL());
    const response = await fetch(browseURL, { method: 'POST' });

    const obj = await response.json() as browseResponse;

    return {
        items: obj?.items,
    };
};