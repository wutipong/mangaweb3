<script lang="ts">
	import { Container, Icon, Button } from 'sveltestrap';
	import { page } from '$app/stores';
	import Item from './Item.svelte';
	import Toolbar from './Toolbar.svelte';
	import type { PageData } from './$types';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import MoveToTop from '$lib/MoveToTop.svelte';
	import { getBackendBaseURL } from '$lib/config';

	export let data: PageData;

	let favoriteOnly = false;
	let aboutDialog: AboutDialog;

	function toggleFavoriteOnly() {
		favoriteOnly = !favoriteOnly;
	}

	function onAboutClick() {
		aboutDialog.show();
	}

	function createThumbnailUrl(name: string) {
		const output = new URL('/tag/thumbnail', getBackendBaseURL());
		output.searchParams.append('tag', name);

		return output;
	}

	function createBrowseURL(name: string) {
		const output = new URL('/browse', $page.url.origin);
		output.searchParams.append('tag', name);

		return output;
	}
</script>

<Toolbar
	title="Tag List"
	browseURL={new URL('/browse', $page.url.origin).toString()}
	tagListURL={$page.url.toString()}
	onFilterFavorite={toggleFavoriteOnly}
	{favoriteOnly}
	{onAboutClick}
/>

<Container fluid style="padding-top:100px;">
	<div class="grid-container">
		{#each data.tags as tag}
			{#if !favoriteOnly || (favoriteOnly && tag.favorite)}
				<Item
					name={tag.name}
					favorite={tag.favorite}
					id={tag.id.toString()}
					url={createBrowseURL(tag.name).toString()}
					thumbnailURL={createThumbnailUrl(tag.name).toString()}
				/>
			{/if}
		{/each}
	</div>
</Container>

<AboutDialog bind:this={aboutDialog} version={'development'} />

<MoveToTop/>