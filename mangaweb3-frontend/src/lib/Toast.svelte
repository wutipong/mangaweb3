<script lang="ts">
	import { Toast, ToastBody, ToastHeader } from 'sveltestrap';

	let message = '';
	let header = '';
	let isOpen = false;

	export function show(h: string, m: string) {
		message = m;
		header = h;

		isOpen = true;
	}
</script>

<div class="toast-container position-fixed bottom-0 end-0 p-3" style="z-index: 11">
	<div id="toast" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
		<div class="toast-header">
			<strong class="me-auto" id="toast-header">{header}</strong>
			<button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close" />
		</div>
		<div class="toast-body" id="toast-content">
			{message}
		</div>
	</div>

	<Toast autohide {isOpen} on:close={() => (isOpen = false)}>
		<ToastHeader>{header}</ToastHeader>
		<ToastBody>{message}</ToastBody>
	</Toast>
</div>
