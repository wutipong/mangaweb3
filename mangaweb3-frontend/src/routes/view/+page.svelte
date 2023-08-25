<script lang="ts">
	import Toolbar from './Toolbar.svelte';
	import ImageViewer from './ImageViewer.svelte';
	//import Toast from "$libs/Toast.svelte";
	// import ModalDialog from "./Common/ModalDialog.svelte";
	import PageScroll from './PageScroll.svelte';
	import type { PageData } from './$types';
	import { getBackendBaseURL } from '$lib/config';

	export let data: PageData;

	let current = 0;
	let viewer: ImageViewer;

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
	/*
    let favorite = params.Favorite;
    let name = params.Name;
    let tags = params.Tags;
    let browseURL = params.BrowseURL;


    let toast;
    let aboutDialog;


    function downloadManga() {
        download(params.DownloadURL);
    }

    function downloadPage() {
        download(params.DownloadPageURLs[current]);
    }

    async function toggleFavorite() {
        favorite = !favorite;

        const urlSearchParams = new URLSearchParams();
        urlSearchParams.set("favorite", favorite.toString());

        const url = new URL(params.SetFavoriteURL, window.location.origin);
        url.search = urlSearchParams.toString();

        await fetch(url);
        if (favorite) {
            toast.show("Favorite", "The current manga is now your favorite.");
        } else {
            toast.show(
                "Favorite",
                "The current manga is no longer your favorite."
            );
        }
    }

    async function updateCover() {
        const url = new URL(
            params.UpdateCoverURLs[current],
            window.location.origin
        );

        await fetch(url);
        toast.show("Update Cover", "The cover image is updated successfully.");
    }

	function download(url) {
        let link = document.createElement("a");
        link.setAttribute("download", "");
        link.href = url;
        document.body.appendChild(link);

        link.click();
        link.remove();
    }

	function onAboutClick() {
        aboutDialog.show();
    }
*/
	function onIndexChange(i: number) {
		current = i;
	}

	function onValueChange(n: number) {
		viewer.advance(n);
	}
</script>

<PageScroll PageCount={data.indices.length} {onValueChange} Current={current}></PageScroll>

<div class="fullscreen" style="padding-top:80px;">
	<ImageViewer imageURLs={createImageUrls(data.name, data.indices)} {onIndexChange} bind:this={viewer} />
</div>

<!-- Toolbar
    Tags={tags}
    Name={name}
    Favorite={favorite}
    BrowseURL={browseURL}
    onDownloadManga={downloadManga}
    onDownloadPage={downloadPage}
    {toggleFavorite}
    {updateCover}
    {onAboutClick}
/ -->

<!-- Toast bind:this={toast} /-->

<!-- ModalDialog Id="aboutModal" Title="About" bind:this={aboutDialog}>
    <h5>MangaWeb</h5>
    <h6>Version {params.Version}</h6>
    <p>&copy; 2021-2023 Wutipong Wongsakuldej. All Right Reserved</p>
    <p>Licensed under MIT License</p>
    <p><a href="https://github.com/wutipong/mangaweb">Homepage</a></p>
</ModalDialog -->
