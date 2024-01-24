<script lang="ts">
	let links = [
		{ name: 'Components', href: '/dev/components' },
		{ name: 'Email', href: '/dev/email' }
	];
</script>

<div>
	<ul>
		{#each links as link}
			<li><a href={link.href}>{link.name}</a></li>
		{/each}
	</ul>
</div>
