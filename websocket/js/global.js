document.addEventListener('DOMContentLoaded', () => {
	document.body.classList.add('loaded');
})

document.addEventListener('scroll', () => {
	window.setTimeout(() => {
		if (window.scrollY > 20) {
			document.body.classList.add('scrolled')
		} else {
			document.body.classList.remove('scrolled')
		}
	}, 250)
})

document.getElementById('menuButton').addEventListener('click', () => {
	document.body.classList.toggle('menuOpen');
})
