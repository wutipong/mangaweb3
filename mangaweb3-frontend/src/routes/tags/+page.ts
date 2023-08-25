import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {

    const tagListURL = new URL("/tag/list", "http://localhost:8972");
    const response = await fetch(tagListURL, { mode: 'no-cors', method: 'POST'});

    const obj = await response.json();
    console.log(obj)
    return {
        post: {
            //tags: resp.tags,
        },
    };
};