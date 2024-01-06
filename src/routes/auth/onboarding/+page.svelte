<script lang="ts">
	import * as Card from '$lib/components/ui/card';

	export let data;

	const { user_profile } = data;
</script>

<Card.Root>
	<Card.Header>{JSON.stringify(user_profile, null, 2)}</Card.Header>
</Card.Root>
