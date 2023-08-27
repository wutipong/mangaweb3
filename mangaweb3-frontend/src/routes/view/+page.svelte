<script lang="ts">
	import { getBackendBaseURL } from '$lib/config';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import ImageViewer from './ImageViewer.svelte';
	import PageScroll from './PageScroll.svelte';
	import Toast from '$lib/Toast.svelte';
	import Toolbar from './Toolbar.svelte';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';

	let current = 0;
	let viewer: ImageViewer;

	interface Request {
		path: string;
	}

	interface Response {
		browse_url: string;
		favorite: boolean;
		indices: number[];
		tags: string[];
	}

	let request: Request = {
		path: ''
	};

	let response: Response = {
		browse_url: '',
		favorite: false,
		indices: [],
		tags: []
	};

	onMount(async () => {
		const params = $page.url.searchParams;
		if (params.has('name')) {
			request.path = params.get('name') as string;
		}

		const url = new URL('/view', getBackendBaseURL());
		const resp = await fetch(url, { method: 'POST', body: JSON.stringify(request) });
		response = await resp.json();
	});

	function createImageUrls(name: string, indices: number[]): string[] {
		const url = new URL('/view/get_image', getBackendBaseURL());
		const output = [];

		url.searchParams.append('name', name);
		for (const i in indices) {
			url.searchParams.set('i', indices[i].toString());
			output.push(url.toString());
		}

		return output;
	}

	let aboutDialog: AboutDialog;
	let toast: Toast;

	function downloadManga() {
		// download(params.DownloadURL);
	}

	function downloadPage() {
		// download(params.DownloadPageURLs[current]);
	}

	async function toggleFavorite() {
		const req = {
			favorite: !response.favorite,
			name: request.path
		};

		const url = new URL('/view/set_favorite', getBackendBaseURL());

		const resp = await fetch(url, { method: 'POST', body: JSON.stringify(req) });
		const json = await resp.json();

		if (json.favorite) {
			toast.show('Favorite', 'The current manga is now your favorite.');
		} else {
			toast.show('Favorite', 'The current manga is no longer your favorite.');
		}

		response.favorite = json.favorite;
	}

	async function updateCover() {
		/*
		const url = new URL(params.UpdateCoverURLs[current], window.location.origin);

		await fetch(url);
		toast.show('Update Cover', 'The cover image is updated successfully.');
        */
	}

	function download(url: string) {
		let link = document.createElement('a');
		link.setAttribute('download', '');
		link.href = url;
		document.body.appendChild(link);

		link.click();
		link.remove();
	}

	function onAboutClick() {
		aboutDialog.show();
	}

	function onIndexChange(i: number) {
		current = i;
	}

	function onValueChange(n: number) {
		viewer.advance(n);
	}
</script>

<PageScroll PageCount={response.indices.length} {onValueChange} Current={current} />

<div class="fullscreen" style="padding-top:80px;">
	<ImageViewer
		imageURLs={createImageUrls(request.path, response.indices)}
		{onIndexChange}
		bind:this={viewer}
	/>
</div>

<Toolbar
	Tags={response.tags}
	Name={request.path}
	Favorite={response.favorite}
	BrowseURL={new URL('/browse', $page.url.origin).toString()}
	onDownloadManga={downloadManga}
	onDownloadPage={downloadPage}
	{toggleFavorite}
	{updateCover}
	{onAboutClick}
/>

<Toast bind:this={toast} />

<AboutDialog bind:this={aboutDialog} />
