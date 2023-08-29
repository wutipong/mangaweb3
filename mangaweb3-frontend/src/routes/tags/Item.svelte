<script lang="ts">
    import { onMount } from 'svelte';
	import { Badge, Card, CardBody, CardFooter, CardImg, Icon } from 'sveltestrap';
	import type {Tag} from '$lib/tag'
	import { page } from '$app/stores'
	import { variables } from '$lib/variables';

    export let tag: Tag

    let classStr = '';

    onMount(async () => {
		if (tag.favorite) {
			classStr = 'border border-pink';
		}
	});

	function createBrowseURL(name: string) {
		const output = new URL('/browse', $page.url.origin);
		output.searchParams.append('tag', name);

		return output;
	}

	function createThumbnailUrl(name: string) {
		const output = new URL('/tag/thumbnail', variables.basePath);
		output.searchParams.append('tag', name);

		return output;
	}
</script>

<Card class="{classStr} h-100" id={tag.id.toString()}>
	<CardImg top src={createThumbnailUrl(tag.name).toString()} />
	<CardBody>
		<a href={createBrowseURL(tag.name).toString()}>{tag.name}</a>
	</CardBody>
	<CardFooter>
		{#if tag.favorite}
			<Badge class="bg-pink">
				<span><Icon name="star-fill" /> Favorite </span>
			</Badge>
        {:else}
        <Badge>
            <span><Icon name="tag" />  </span>
        </Badge>
		{/if}
	</CardFooter>
</Card>