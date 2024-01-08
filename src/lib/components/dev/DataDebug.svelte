<script lang="ts">
	export let data: unknown;
	export let label: string | null = null;
	export let ref = undefined;

	function syntaxHighlight(json: unknown) {
		switch (typeof json) {
			case 'function': {
				return `<span class="function">[function ${json.name ?? 'unnamed'}]</span>`;
			}
			case 'symbol': {
				return `<span class="symbol">${json.toString()}</span>`;
			}
		}

		const encodedString = JSON.stringify(
			json,
			function (key, value) {
				if (value === undefined) {
					return '#}#undefined';
				}
				if (typeof this === 'object' && this[key] instanceof Date) {
					return '#}D#' + (isNaN(this[key]) ? 'Invalid Date' : value);
				}
				if (typeof value === 'number' && isNaN(value)) {
					return '#}#NaN';
				}
				if (typeof value === 'bigint') {
					return '#}BI#' + value;
				}
				if (value instanceof Error) {
					return '#}E#' + `${value.name}: ${value.message || value.cause || '(No error message)'}`;
				}
				return value;
			},
			2
		)
			.replace(/&/g, '&amp;')
			.replace(/</g, '&lt;')
			.replace(/>/g, '&gt;');

		return encodedString.replace(
			/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+-]?\d+)?)/g,
			function (match) {
				let cls = 'number';
				if (/^"/.test(match)) {
					if (/:$/.test(match)) {
						cls = 'key';
						match = match.slice(1, -2) + ':';
					} else {
						cls = 'string';
						if (match == '"#}#NaN"') {
							cls = 'nan';
							match = 'NaN';
						} else if (match == '"#}#undefined"') {
							cls = 'undefined';
							match = 'undefined';
						} else if (match.startsWith('"#}D#')) {
							cls = 'date';
							match = match.slice(5, -1);
						} else if (match.startsWith('"#}BI#')) {
							cls = 'bigint';
							match = match.slice(6, -1) + 'n';
						} else if (match.startsWith('"#}F#')) {
							cls = 'function';
							match = match.slice(5, -1);
						} else if (match.startsWith('"#}E#')) {
							cls = 'error';
							match = match.slice(5, -1);
						}
					}
				} else if (/true|false/.test(match)) {
					cls = 'boolean';
				} else if (/null/.test(match)) {
					cls = 'null';
				}
				return '<span class="' + cls + '">' + match + '</span>';
			}
		);
	}

	$: debugData = data;
</script>

<div class="super-debug">
	{#if label}
		<div class="super-debug--label px-4 py-1 text-xl">{label}</div>
	{/if}
	<pre class="super-debug--pre" bind:this={ref}>
    <code class="super-debug--code">
	    {@html syntaxHighlight(debugData)}
    </code>
  </pre>
</div>

<style>
	.super-debug {
		--_sd-bg-color: var(--sd-bg-color, var(--sd-vscode-bg-color, rgb(30, 41, 59)));
		position: relative;
		background-color: var(--_sd-bg-color);
		border-radius: 0.5rem;
		overflow: hidden;
	}

	.super-debug pre {
		color: var(--sd-code-default, var(--sd-vscode-code-default, #999));
		background-color: var(--_sd-bg-color);
		font-size: 1em;
		margin-bottom: 0;
	}

	:global(.super-debug--code .key) {
		color: var(--sd-code-key, var(--sd-vscode-code-key, #eab308));
	}

	:global(.super-debug--code .string) {
		color: var(--sd-code-string, var(--sd-vscode-code-string, #6ec687));
	}

	:global(.super-debug--code .date) {
		color: var(--sd-code-date, var(--sd-vscode-code-date, #f06962));
	}

	:global(.super-debug--code .boolean) {
		color: var(--sd-code-boolean, var(--sd-vscode-code-boolean, #79b8ff));
	}

	:global(.super-debug--code .number) {
		color: var(--sd-code-number, var(--sd-vscode-code-number, #af77e9));
	}

	:global(.super-debug--code .bigint) {
		color: var(--sd-code-bigint, var(--sd-vscode-code-bigint, #af77e9));
	}

	:global(.super-debug--code .null) {
		color: var(--sd-code-null, var(--sd-vscode-code-null, #238afe));
	}

	:global(.super-debug--code .nan) {
		color: var(--sd-code-nan, var(--sd-vscode-code-nan, #af77e9));
	}

	:global(.super-debug--code .undefined) {
		color: var(--sd-code-undefined, var(--sd-vscode-code-undefined, #238afe));
	}

	:global(.super-debug--code .function) {
		color: var(--sd-code-function, var(--sd-vscode-code-function, #f06962));
	}

	:global(.super-debug--code .symbol) {
		color: var(--sd-code-symbol, var(--sd-vscode-code-symbol, #4de0c5));
	}

	:global(.super-debug--code .error) {
		color: var(--sd-code-error, var(--sd-vscode-code-error, #ff475d));
	}

	.super-debug--label {
		color: var(--sd-label-color, var(--sd-vscode-label-color, white));
	}

	.super-debug pre::-webkit-scrollbar {
		width: var(--sd-sb-width, var(--sd-vscode-sb-width, 1.25rem));
		height: var(--sd-sb-height, var(--sd-vscode-sb-height, 1.25rem));
	}

	.super-debug pre::-webkit-scrollbar-track {
		border-radius: 12px;
		background-color: var(
			--sd-sb-track-color,
			var(--sd-vscode-sb-track-color, hsl(0, 0%, 40%, 0.2))
		);
	}
	.super-debug:is(:focus-within, :hover) pre::-webkit-scrollbar-track {
		border-radius: 12px;
		background-color: var(
			--sd-sb-track-color-focus,
			var(--sd-vscode-sb-track-color-focus, hsl(0, 0%, 50%, 0.2))
		);
	}

	.super-debug pre::-webkit-scrollbar-thumb {
		border-radius: 12px;
		background-color: var(
			--sd-sb-thumb-color,
			var(--sd-vscode-sb-thumb-color, hsl(217, 50%, 50%, 0.5))
		);
	}
	.super-debug:is(:focus-within, :hover) pre::-webkit-scrollbar-thumb {
		border-radius: 12px;
		background-color: var(
			--sd-sb-thumb-color-focus,
			var(--sd-vscode-sb-thumb-color-focus, hsl(217, 50%, 50%))
		);
	}
</style>
