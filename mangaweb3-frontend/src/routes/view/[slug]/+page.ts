import type { LoadEvent } from '@sveltejs/kit';

export function load({ params }: LoadEvent) {
    return {path: params.slug}
}