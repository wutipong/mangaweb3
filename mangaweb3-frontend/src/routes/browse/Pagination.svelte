<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { Pagination, PaginationItem, PaginationLink } from 'sveltestrap';

	export let currentPage = 0;
	export let totalPage = 1;
	export let pageToShow = 5;
	export let onPageClick = (i: number): void => {};

	interface Page {
		url: URL;
		index: number;
	}

	let numberUrls: Page[] = [];
	let firstUrl: URL;
	let lastUrl: URL;

	onMount(() => {
		if (totalPage == 0) {
			return;
		}

		firstUrl = new URL($page.url);
		firstUrl.searchParams.set('page', '0');

		lastUrl = new URL($page.url);
		lastUrl.searchParams.set('page', (totalPage - 1).toString());

		let halfCount = Math.floor(pageToShow / 2);

		for (let i = currentPage - halfCount; i < currentPage + halfCount; i++) {
			if (i < 0) continue;
			if (i >= totalPage) continue;

			const url = new URL($page.url);
			url.searchParams.set('page', `${i}`);

			numberUrls = [
				...numberUrls,
				{
					url: url,
					index: i,
				}
			];
		}
	});
</script>

<Pagination>
	<PaginationItem>
		<PaginationLink first on:click={() => onPageClick(0)} href={firstUrl?.toString()} />
	</PaginationItem>

	{#each numberUrls as u}
		<PaginationItem>
			<PaginationLink
				on:click={() => {
					onPageClick(u.index);
				}}
				href={u.url.toString()}>{u.index}</PaginationLink
			>
		</PaginationItem>
	{/each}

	<PaginationItem>
		<PaginationLink last on:click={() => onPageClick(totalPage - 1)} href={lastUrl?.toString()} />
	</PaginationItem>
</Pagination>
