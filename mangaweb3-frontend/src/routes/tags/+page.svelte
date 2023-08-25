<script lang="ts">
	import Toolbar from '$lib/tags/Toolbar.svelte';
	import Item from '$lib/tags/Item.svelte';
	import type AboutDialog from '$lib/AboutDialog.svelte';
	import type { PageData } from './$types';
    import { page } from '$app/stores';
    import { Icon } from 'sveltestrap';

	export let data: PageData;

	let favoriteOnly = false;
	let aboutDialog: AboutDialog;

	function toggleFavoriteOnly() {
		favoriteOnly = !favoriteOnly;
	}

	function onAboutClick() {
		aboutDialog.show();
	}

	function CreateThumbnailUrl(name: string) {
		const output = new URL('/tag/thumbnail', 'http://localhost:8972');
		output.searchParams.append('tag', name);

		return output;
	}

    function CreateBrowseURL(name: string) {
        const output = new URL('/browse', $page.url.origin);
		output.searchParams.append('tag', name);

		return output;
    }
</script>

<!-- Toolbar
    title={params.Title}
    browseURL={params.BrowseURL}
    tagListURL={params.TagListURL}
    onFilterFavorite={toggleFavoriteOnly}
    {favoriteOnly}
    {onAboutClick}
/-->

<div class="container-fluid" style="padding-top:100px;">
	<div class="grid-container">
		{#each data.tags as tag}
			{#if !favoriteOnly || (favoriteOnly && tag.favorite)}
				<Item
					name={tag.name}
					favorite={tag.favorite}
					id={tag.id.toString()}
					url={CreateBrowseURL(tag.name).toString()}
					thumbnailURL={CreateThumbnailUrl(tag.name).toString()}
				/>
			{/if}
		{/each}
	</div>
</div>

<!-- AboutDialog bind:this={aboutDialog} version={params.Version} /-->

<nav aria-label="Move to top navigation" class="position-fixed bottom-0 end-0 p-3">
	<a class="btn btn-secondary" href="#top">
		<i class="bi bi-chevron-double-up" />
		<span class="d-none d-sm-block">Top</span>
	</a>
</nav>
