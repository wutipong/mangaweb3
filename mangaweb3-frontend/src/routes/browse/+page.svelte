<script lang="ts">
	import { getBackendBaseURL } from '$lib/config';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import {
		Spinner,
		Icon,
		Collapse,
		Navbar,
		NavbarToggler,
		NavbarBrand,
		Nav,
		NavItem,
		NavLink,
		Dropdown,
		DropdownToggle,
		DropdownMenu,
		DropdownItem,
		InputGroup,
		InputGroupText,
		Input,
		Button
	} from 'sveltestrap';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import Item from './Item.svelte';
	import MoveToTop from '$lib/MoveToTop.svelte';
	import Pagination from './Pagination.svelte';
	import Toast from '$lib/Toast.svelte';
	import Toolbar from './Toolbar.svelte';
	import { goto } from '$app/navigation';

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

	let request = defaultRequest();

	let response: Response = {
		items: [],
		pages: [],
		tag_favorite: true,
		total_page: 0
	};

	let promise: Promise<void>;
	onMount(() => {
		parseParams();
		invalidate(request);
	});

	function defaultRequest(): Request {
		return {
			favorite_only: false,
			item_per_page: 30,
			order: 'descending',
			page: 0,
			search: '',
			sort: 'createTime',
			tag: ''
		};
	}

	function parseParams() {
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

		if (params.has('page')) {
			let v = params.get('page');
			if (v != null) {
				request.page = parseInt(v);
			}
		}
	}

	async function loadData() {
		let u = new URL('/browse', getBackendBaseURL());
		const r = await fetch(u, { method: 'POST', body: JSON.stringify(request) });
		response = await r.json();
	}

	function changeSort(sortBy: 'name' | 'createTime') {
		request.sort = sortBy;
		switch (sortBy) {
			case 'name':
				request.order = 'ascending';
			case 'createTime':
				request.order = 'descending';
		}

		request.page = 0;
		invalidate(request);
	}

	function changeOrder(order: 'ascending' | 'descending') {
		request.order = order;
		request.page = 0;

		invalidate(request);
	}

	function onFilterFavorite() {
		let url = $page.url;
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
		const req = {
			favorite: !response.tag_favorite,
			tag: request.tag
		};

		const url = new URL('/tag/set_favorite', getBackendBaseURL());

		const resp = await fetch(url, { method: 'POST', body: JSON.stringify(req) });
		const json = await resp.json();

		if (json.favorite) {
			toast.show('Favorite', `The tag "${request.tag}" is now your favorite.`);
		} else {
			toast.show('Favorite', `The tag "${request.tag}" is no longer your favorite.`);
		}

		response.tag_favorite = json.favorite;
	}

	let searchText = '';
	function onSearchClick() {
		request.search = searchText;
		invalidate(request);
	}

	function onPageClick(i: number): void {
		request.page = i;
		invalidate(request);
	}

	let navbarToggleOpen = false;
	function handleUpdate(event: CustomEvent<boolean>) {
		navbarToggleOpen = event.detail;
	}

	function invalidate(req: Request) {
		request = req;

		$page.url.searchParams.set('favorite_only', `${request.favorite_only}`);
		$page.url.searchParams.set('order', request.order);
		$page.url.searchParams.set('page', request.page.toString());

		if (request.search == '') {
			$page.url.searchParams.delete('search');
		} else {
			$page.url.searchParams.set('search', request.tag);
		}

		$page.url.searchParams.set('sort', request.sort);

		if (request.tag == '') {
			$page.url.searchParams.delete('tag');
		} else {
			$page.url.searchParams.set('tag', request.tag);
		}

		goto($page.url);
		promise = loadData();
	}
</script>

{#await promise}
	<div><Spinner type="grow" /> Loading ...</div>
{:then}
	<Navbar color="dark" dark expand="md" sticky={'top'}>
		<NavbarBrand href="/">{request.tag == '' ? 'Browse' : `Browse: ${request.tag}`}</NavbarBrand>
		<NavbarToggler on:click={() => (navbarToggleOpen = !navbarToggleOpen)} />
		<Collapse isOpen={navbarToggleOpen} navbar expand="md" on:update={handleUpdate}>
			<Nav navbar>
				<Dropdown nav inNavbar>
					<DropdownToggle nav caret>Browse</DropdownToggle>
					<DropdownMenu end>
						<DropdownItem on:click={()=> {request = defaultRequest(); invalidate(request)}}>
							<Icon name="list-ul" class="me-3" /> All items
						</DropdownItem>
						<DropdownItem on:click={() => goto(new URL('/tags', $page.url.origin).toString())}>
							<Icon name="tags-fill" class="me-3" /> Tag list
						</DropdownItem>
					</DropdownMenu>
				</Dropdown>
				<Dropdown nav inNavbar>
					<DropdownToggle nav caret>Sort By</DropdownToggle>
					<DropdownMenu end>
						<DropdownItem active={request.sort == 'name'} on:click={() => changeSort('name')}>
							<Icon name="type" class="me-3" /> Name
						</DropdownItem>
						<DropdownItem
							active={request.sort == 'createTime'}
							on:click={() => changeSort('createTime')}
						>
							<Icon name="clock" class="me-3" /> Create time
						</DropdownItem>
						<DropdownItem divider />
						<DropdownItem
							active={request.order == 'ascending'}
							on:click={() => changeOrder('ascending')}
						>
							<Icon name="sort-down-alt" class="me-3" />Ascending
						</DropdownItem>
						<DropdownItem
							active={request.order == 'descending'}
							on:click={() => changeOrder('descending')}
						>
							<Icon name="sort-up-alt" class="me-3" /> Descending
						</DropdownItem>
					</DropdownMenu>
				</Dropdown>
				<NavItem>
					<NavLink on:click={() => aboutDialog.show()}>About</NavLink>
				</NavItem>
			</Nav>
			<Nav class="ms-auto" navbar>
				<NavItem hidden={request.tag != ''}>
					<Button on:click={() => onTagFavorite()}>
						<Icon name="star-fill" class="me-3" /> Favorite tag
					</Button>
				</NavItem>
			</Nav>
			<Nav navbar>
				<NavItem>
					<InputGroup>
						<Input type="text" bind:value={request.search} />
						<Button on:click={() => onSearchClick()}>Search</Button>
					</InputGroup>
				</NavItem>
			</Nav>
		</Collapse>
	</Navbar>

	<div class="container-fluid" style="padding-top:30px;">
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

	<div aria-label="Page navigation" class="position-fixed bottom-0 start-50 p-3 translate-middle-x">
		<Pagination currentPage={request.page} totalPage={response.total_page} {onPageClick} />
	</div>
{:catch}
	<Icon name="exclamation-octagon-fill" color="danger" /> Cannot fetch browse data.
{/await}

<Toast bind:this={toast} />

<MoveToTop />

<AboutDialog bind:this={aboutDialog} />
