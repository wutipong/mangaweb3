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

	let promise: Promise<void>;
	onMount(() => {
		promise = loadData();
	});

	async function loadData() {
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

		let u = new URL('/browse', getBackendBaseURL());
		const r = await fetch(u, { method: 'POST', body: JSON.stringify(request) });
		response = await r.json();
	}

	function changeSort(sortBy: string) {
		let url = $page.url;
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
		let url = $page.url;
		let searchParams = new URLSearchParams(url.search);

		searchParams.set('order', order);

		url.search = searchParams.toString();
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

	function onSearchClick(t: string) {
		let searchText = t;
		let url = $page.url;
		let searchParams = new URLSearchParams(url.search);
		searchParams.set('search', searchText);

		url.search = searchParams.toString();
	}

	function onAboutClick() {
		aboutDialog.show();
	}

	function onPageClick(i: number): void {
		const params = $page.url.searchParams;
		if ($page.url.searchParams.has('page')) {
			$page.url.searchParams.delete('page');
		}

		$page.url.searchParams.set('page', i.toString());
		promise = loadData();
	}

	let navbarToggleOpen = false;
	function handleUpdate(event: CustomEvent<boolean>) {
		navbarToggleOpen = event.detail;
	}
</script>

{#await promise}
	<div><Spinner type="grow" /> Loading ...</div>
{:then}
	<Navbar color="dark" dark expand="md" sticky={'top'}>
		<NavbarBrand href="/">{request.tag == '' ? 'Browse' : `Browse ${request.tag}`}</NavbarBrand>
		<NavbarToggler on:click={() => (navbarToggleOpen = !navbarToggleOpen)} />
		<Collapse isOpen={navbarToggleOpen} navbar expand="md" on:update={handleUpdate}>
			<Nav navbar>
				<Dropdown nav inNavbar>
					<DropdownToggle nav caret>Browse</DropdownToggle>
					<DropdownMenu end>
						<DropdownItem><Icon name="list-ul" class="me-3" /> All items</DropdownItem>
						<DropdownItem><Icon name="tags-fill" class="me-3" /> Tag list</DropdownItem>
					</DropdownMenu>
				</Dropdown>
				<Dropdown nav inNavbar>
					<DropdownToggle nav caret>Sort By</DropdownToggle>
					<DropdownMenu end>
						<DropdownItem><Icon name="type" class="me-3" /> Name</DropdownItem>
						<DropdownItem><Icon name="clock" class="me-3" /> Create Time</DropdownItem>
						<DropdownItem divider />
						<DropdownItem><Icon name="sort-down-alt" class="me-3" />Ascending</DropdownItem>
						<DropdownItem><Icon name="sort-up-alt" class="me-3" /> Descending</DropdownItem>
					</DropdownMenu>
				</Dropdown>
				<NavItem>
					<NavLink on:click={()=>aboutDialog.show()}>About</NavLink>
				</NavItem>
			</Nav>
			<Nav class="ms-auto" navbar>
				<NavItem>
					<Button>
						<Icon name="star-fill" class=" me-3" /> Favorite tag
					</Button>
				</NavItem>
			</Nav>
			<Nav navbar>
				<NavItem>
					<InputGroup>
						<Input type="text" />
						<Button>Search</Button>
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
