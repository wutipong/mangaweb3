<script lang="ts">
	import { page } from '$app/stores';
	export let tags: string[] = [];
	export let name = '';
	export let favorite = false;
	export let browseURL = '';

	export let onDownloadManga = (): void => {};
	export let onDownloadPage = (): void => {};
	export let toggleFavorite = (): void => {};
	export let updateCover = (): void => {};
	export let onAboutClick = (): void => {};

	function createBrowseTagURL(tag: string): URL {
		const u = new URL('/browse', $page.url.origin);
		u.searchParams.set('tag', tag);

		return u;
	}
</script>

<nav class="navbar navbar-dark bg-dark fixed-top navbar-expand-lg">
	<div class="container-fluid">
		<span class="navbar-brand mr-auto text-truncate h1 d-none d-sm-block" style="max-width: 50%;"
			>{name}</span
		>
		<span class="navbar-brand mr-auto text-truncate h1 d-block d-sm-none" style="max-width: 50%;"
			>View</span
		>

		<button
			class="navbar-toggler"
			type="button"
			data-bs-toggle="collapse"
			data-bs-target="#navbarSupportedContent"
			aria-controls="navbarSupportedContent"
			aria-expanded="false"
			aria-label="Toggle navigation"
		>
			<span class="navbar-toggler-icon" />
		</button>

		<div class="collapse navbar-collapse" id="navbarSupportedContent">
			<ul class="navbar-nav mb-2 me-auto">
				<li class="nav-item dropdown">
					<button
						class="nav-link dropdown-toggle"
						id="navbarDropdown"
						data-bs-toggle="dropdown"
						aria-haspopup="true"
						aria-expanded="false"
					>
						Tags
					</button>
					<div class="dropdown-menu" aria-labelledby="navbarDropdown">
						{#each tags as tag}
							<a class="dropdown-item" href={createBrowseTagURL(tag).toString()}>
								<i class="bi bi-tag me-3" />{tag}
							</a>
						{/each}
					</div>
				</li>
				<li class="nav-item dropdown">
					<button
						class="nav-link dropdown-toggle"
						id="navbarDropdown"
						data-bs-toggle="dropdown"
						aria-haspopup="true"
						aria-expanded="false"
					>
						Tools
					</button>
					<div class="dropdown-menu" aria-labelledby="navbarDropdown">
						<button class="dropdown-item" id="download-btn" on:click={onDownloadPage}>
							<i class="bi bi-download me-3" />
							Download Current Page
						</button>
						<button class="dropdown-item" on:click={onDownloadManga}>
							<i class="bi bi-download me-3" />
							Download Manga
						</button>
						<div class="dropdown-divider" />
						<button class="dropdown-item" id="update-cover-btn" on:click={updateCover}>
							<i class="bi bi-journal-arrow-up me-3" />
							Update Cover
						</button>
					</div>
				</li>
				<li class="nav-item">
					<button class="nav-link" on:click={onAboutClick}>About</button>
				</li>
			</ul>
			<ul class="navbar-nav">
				<li class="nav-item mb-2 mb-lg-0 ms-lg-2">
					<button
						id="favorite-btn"
						class="btn"
						class:btn-pink={favorite}
						class:active={favorite}
						class:btn-outline-pink={!favorite}
						on:click={toggleFavorite}
					>
						<i class="bi bi-star-fill me-3" /> Favorite
					</button>
				</li>
				<li class="nav-item mb-2 mb-lg-0 ms-lg-2">
					<a id="close-btn" class="btn btn-danger" href={browseURL}>
						<i class="bi bi-x-circle-fill me-3" /> Close
					</a>
				</li>
			</ul>
		</div>
	</div>
</nav>
