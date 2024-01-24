<script lang="ts">
	import * as Card from '$lib/components/ui/card';
</script>

<p class="mx-auto w-full py-10 text-center text-4xl">Welcome!</p>

<Card.Root>
	<Card.Content class="pt-4">
		<slot />
	</Card.Content>
</Card.Root>
