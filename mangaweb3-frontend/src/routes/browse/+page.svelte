<script lang="ts">
	import Toolbar from './Toolbar.svelte';
	import Item from './Item.svelte';
	// import Pagination from "./Common/Pagination.svelte";
	// import PageItem from "./Common/PageItem.svelte";
	import Toast from '$lib/Toast.svelte';
	import MoveToTop from '$lib/MoveToTop.svelte';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { getBackendBaseURL } from '$lib/config';

	interface Request {
		favorite_only: boolean;
		item_per_page: number;
		order: 'ascending' | 'descending';
		page: number;
		search: string;
		sort: 'name' | 'createTime';
		tag: string;
	}

	interface Item {
		create_time: string;
		favorite: boolean;
		id: number;
		is_read: boolean;
		name: string;
	}

	interface Page {
		content: string;
		is_active: boolean;
		is_enabled: boolean;
		is_hidden_on_small: boolean;
		link_url: string;
	}

	interface Response {
		items: Item[];
		pages: Page[];
		tag_favorite: boolean;
		total_page: number;
	}

	let toast: Toast;
	let aboutDialog: AboutDialog;

	let request: Request = {
		favorite_only: false,
		item_per_page: 30,
		order: 'descending',
		page: 0,
		search: '',
		sort: 'createTime',
		tag: ''
	};
	let response: Response = {
		items: [],
		pages: [],
		tag_favorite: true,
		total_page: 0
	};

	onMount(async () => {
		const params = $page.url.searchParams;
		if (params.has('favorite_only')) {
			request.favorite_only = params.get('favorite_only') == 'true';
		}

		if (params.has('order')) {
			const v = params.get('order');
			if (v == 'ascending') {
				request.order = 'ascending';
			} else if (v == 'descending') {
				request.order = 'descending';
			}
		}

		if (params.has('tag')) {
			request.tag = params.get('tag') as string;
		}

		let u = new URL('/browse', getBackendBaseURL());
		const r = await fetch(u, { method: 'POST', body: JSON.stringify(request) });
		response = await r.json();
	});

	function changeSort(sortBy: string) {
		let url = window.location;
		let searchParams = new URLSearchParams(url.search);
		searchParams.set('sort', sortBy);

		if (sortBy === 'name') {
			searchParams.set('order', 'ascending');
		} else if (sortBy === 'createTime') {
			searchParams.set('order', 'descending');
		}

		searchParams.delete('page');

		url.search = searchParams.toString();
	}

	function changeOrder(order: string) {
		let url = window.location;
		let searchParams = new URLSearchParams(url.search);

		searchParams.set('order', order);

		url.search = searchParams.toString();
	}

	function onFilterFavorite() {
		let url = window.location;
		let searchParams = new URLSearchParams(url.search);

		let isFavorite = request.favorite_only;
		searchParams.set('favorite', (!isFavorite).toString());

		url.search = searchParams.toString();
	}

	async function rescanLibrary() {
		const url = new URL('/browse/rescan_library', getBackendBaseURL());
		await fetch(url);
		toast.show(
			'Re-scan Library',
			'Library re-scanning in progress. Please refresh after a few minutes.'
		);
	}

	async function onTagFavorite() {
		const urlSearchParams = new URLSearchParams();
		urlSearchParams.set('favorite', response.tag_favorite.toString());

		const req = {
			favorite: !response.tag_favorite,
			tag: request.tag
		};

		const url = new URL('/tag/set_favorite', window.location.origin);

		const resp = await fetch(url, { method: 'POST', body: JSON.stringify(req) });
		const json = await resp.json();

		if (json.favorite) {
			toast.show('Favorite', `The tag "${request.tag}" is now your favorite.`);
		} else {
			toast.show('Favorite', `The tag "${request.tag}" is no longer your favorite.`);
		}

		response.tag_favorite = json.favorite;
	}

	function onSearchClick(t: string) {
		let searchText = t;
		let url = window.location;
		let searchParams = new URLSearchParams(url.search);
		searchParams.set('search', searchText);

		url.search = searchParams.toString();
	}

	function onAboutClick() {
		aboutDialog.show();
	}
</script>

<Toolbar
	Title={request.tag == '' ? 'Browse' : `Browse ${request.tag}`}
	BrowseURL={new URL('/browse', $page.url.origin).toString()}
	TagListURL={new URL('/tags', $page.url.origin).toString()}
	SortBy={request.sort}
	SortOrder={request.order}
	FavoriteOnly={request.favorite_only}
	Tag={request.tag}
	TagFavorite={response.tag_favorite}
	{changeSort}
	{changeOrder}
	{onFilterFavorite}
	{rescanLibrary}
	{onTagFavorite}
	{onSearchClick}
	SearchText={request.search}
	{onAboutClick}
/>

<div class="container-fluid" style="padding-top:100px;">
	<div class="grid-container">
		{#each response.items as item}
			<Item
				favorite={item.favorite}
				isRead={item.is_read}
				id={item.id.toString()}
				name={item.name}
			/>
		{/each}
	</div>
</div>
<div style="height: 100px;" />

<!-- Pagination>
    {#each params.Pages as page}
        <PageItem
            IsActive={page.IsActive}
            IsEnabled={page.IsEnabled}
            IsHiddenOnSmall={page.IsHiddenOnSmall}
            URL={page.LinkURL}
            Content={page.Content}
        />
    {/each}
</Pagination -->

<Toast bind:this={toast} />

<MoveToTop />

<AboutDialog bind:this={aboutDialog} />
