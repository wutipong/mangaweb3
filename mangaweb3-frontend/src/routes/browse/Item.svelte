<script lang="ts">
	import { getBackendBaseURL } from '$lib/config';
	import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { Card, CardBody } from 'sveltestrap';

	export let favorite = false;
	export let isRead = false;
	export let id = '';
	export let name = '';

    let thumbnailURL = '';
    let viewURL = '';

	onMount(async () => {
        let u = new URL("/browse/thumbnail", getBackendBaseURL());
        u.searchParams.append('name', name);
        thumbnailURL = u.toString();

        u = new URL("/view", $page.url.origin)
        u.searchParams.append('name', name);
        viewURL = u.toString();
		
	});
</script>

<div
	class="card mb-3"
	class:border={favorite || !isRead}
	class:border-pink={favorite}
	class:border-yellow={!isRead}
	{id}
>
	{#if favorite}
		<span class="position-absolute bottom-0 end-0 p-2 badge bg-pink">
			<span><i class="bi bi-star-fill" /> Favorite </span>
		</span>
	{/if}

	{#if !isRead}
		<span class="position-absolute bottom-0 end-0 p-2 badge bg-yellow">
			<span><i class="bi bi-lightning-fill" /> New </span>
		</span>
	{/if}

	<div class="row g-0">
		<div class="col-md-4">
			<a href={viewURL} class="thumb-link">
				<img
					class="img-fluid lazy rounded thumb-img"
					loading="lazy"
					src={thumbnailURL}
					alt={name}
				/>
			</a>
		</div>
		<div class="col-md-8">
			<div class="card-body">
				<h5 class="card-title">
					<a href={viewURL}>{name}</a>
				</h5>
			</div>
		</div>
	</div>
</div>
