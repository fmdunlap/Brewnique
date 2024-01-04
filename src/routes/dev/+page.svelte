<script lang="ts">
	let links = [
		{ name: 'Components', href: '/dev/components' },
		{ name: 'Layout', href: '/dev/layout' },
		{ name: 'Page', href: '/dev/page' },
		{ name: 'Supabase', href: '/dev/supabase' }
	];
</script>

<div>
	<ul>
		{#each links as link}
			<li><a href={link.href}>{link.name}</a></li>
		{/each}
	</ul>
</div>
