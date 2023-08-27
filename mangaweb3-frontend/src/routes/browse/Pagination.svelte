<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte/types/runtime/internal/lifecycle';
	import { Pagination, PaginationItem, PaginationLink } from 'sveltestrap';

	export let currentPage = 0;
	export let pageCount = 1;
	export let pageToShow = 5;

	let firstUrl = '';
	let lastUrl = '';

	interface NumberURL {
		page: string;
		url: string;
	}
	let numberUrls: NumberURL[] = [];

	onMount(() => {
		let url = $page.url;
		let first = new URL(url);
		first.searchParams.delete('page');
		first.searchParams.append('page', '0');
		firstUrl = first.toString();

		let last = new URL(url);
		last.searchParams.delete('page');
		last.searchParams.append('page', `${pageCount - 1}`);
		lastUrl = last.toString();

		let halfCount = pageToShow / 2;

		for (let i = currentPage - halfCount; i < currentPage + halfCount; i++) {
			if (i < 0) continue;
			if (i >= pageCount) continue;

			let u = new URL(url);
			u.searchParams.delete('page');
			u.searchParams.append('page', `${i}`);
			numberUrls.push({
				page: `${i}`,
				url: u.toString()
			});
		}
	});
</script>

<Pagination>
	<PaginationItem>
		<PaginationLink first href={firstUrl} />
	</PaginationItem>

	{#each numberUrls as u}
		<PaginationItem>
			<PaginationLink href={u.url}>{u.page}</PaginationLink>
		</PaginationItem>
	{/each}

	<PaginationItem>
		<PaginationLink last href={lastUrl} />
	</PaginationItem>
</Pagination>
