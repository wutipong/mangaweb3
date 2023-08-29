<script lang="ts">
	import { page } from '$app/stores';
	import { Pagination, PaginationItem, PaginationLink } from 'sveltestrap';

	export let currentPage = 0;
	export let totalPage = 1;
	export let pageToShow = 5;

	interface Page {
		url: URL;
		index: number;
	};

	let numberUrls:Page[] = [];
	let firstUrl: URL;
	let lastUrl: URL;

	$: if(totalPage != 0) {
		firstUrl = new URL($page.url);
		firstUrl.searchParams.set('page', '0');

		lastUrl = new URL($page.url);
		lastUrl.searchParams.set('page', (totalPage - 1).toString());

		let halfCount = Math.floor(pageToShow / 2);

		const startPage = currentPage - halfCount;
		const endPage = startPage + pageToShow;

		let items: Page[] = []
		for (let i = startPage; i < endPage; i++) {
			if (i < 0 || i >= totalPage) continue;

			const url = new URL($page.url);
			url.searchParams.set('page', `${i}`);

			items = [
				...items,
				{
					url: url,
					index: i
				}
			];
		}
		numberUrls = items
	};
</script>

<Pagination>
	<PaginationItem>
		<PaginationLink first href={firstUrl?.toString()} />
	</PaginationItem>

	{#each numberUrls as u}
		<PaginationItem active={u.index == currentPage}>
			<PaginationLink href={u.url.toString()}>
				{u.index}
			</PaginationLink>
		</PaginationItem>
	{/each}

	<PaginationItem>
		<PaginationLink last href={lastUrl?.toString()} />
	</PaginationItem>
</Pagination>
