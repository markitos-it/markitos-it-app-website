// Generate Table of Contents
function generateTableOfContents() {
    const content = document.querySelector('.doc-content');
    const tocContainer = document.getElementById('tableOfContents');

    if (!content || !tocContainer) return;

    const headings = content.querySelectorAll('h2, h3');

    if (headings.length === 0) {
        tocContainer.innerHTML = '<p style="color: var(--text-light); font-size: 0.85rem;">No headings found</p>';
        return;
    }

    const usedIds = new Set();
    const slugify = (value) => value
        .toLowerCase()
        .normalize('NFD').replace(/\p{Diacritic}/gu, '')
        .replace(/[^a-z0-9\s-]/g, '')
        .trim()
        .replace(/\s+/g, '-')
        .replace(/-+/g, '-');

    headings.forEach((heading, index) => {
        const level = heading.tagName.toLowerCase();
        const text = heading.textContent.trim();

        let id = heading.id;
        if (!id) {
            id = slugify(text) || `heading-${index}`;
        }

        let uniqueId = id;
        let counter = 2;
        while (usedIds.has(uniqueId)) {
            uniqueId = `${id}-${counter}`;
            counter += 1;
        }
        usedIds.add(uniqueId);
        heading.id = uniqueId;

        // Create TOC link
        const link = document.createElement('a');
        link.href = `#${uniqueId}`;
        link.textContent = text;
        link.dataset.level = level.replace('h', '');

        link.addEventListener('click', (e) => {
            e.preventDefault();
            heading.scrollIntoView({ behavior: 'smooth', block: 'start' });

            // Update URL without scrolling
            history.pushState(null, null, `#${uniqueId}`);

            // Update active state
            updateActiveTocLink(link);
        });

        tocContainer.appendChild(link);
    });
}

// Update active TOC link
function updateActiveTocLink(activeLink) {
    document.querySelectorAll('.doc-toc a').forEach(link => {
        link.classList.remove('active');
    });
    activeLink.classList.add('active');
}

// Scroll spy for TOC
function initScrollSpy() {
    const headings = document.querySelectorAll('.doc-content h2, .doc-content h3');
    const tocLinks = document.querySelectorAll('.doc-toc a');

    if (headings.length === 0 || tocLinks.length === 0) return;

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const id = entry.target.id;
                const activeLink = document.querySelector(`.doc-toc a[href="#${id}"]`);
                if (activeLink) {
                    updateActiveTocLink(activeLink);
                }
            }
        });
    }, {
        rootMargin: '-80px 0px -80% 0px'
    });

    headings.forEach(heading => observer.observe(heading));
}

// Share functions
function shareDocument(platform) {
    const url = encodeURIComponent(window.location.href);
    const title = encodeURIComponent(document.querySelector('.doc-main-title').textContent);

    let shareUrl;

    switch (platform) {
        case 'twitter':
            shareUrl = `https://twitter.com/intent/tweet?url=${url}&text=${title}`;
            break;
        case 'linkedin':
            shareUrl = `https://www.linkedin.com/sharing/share-offsite/?url=${url}`;
            break;
        case 'facebook':
            shareUrl = `https://www.facebook.com/sharer/sharer.php?u=${url}`;
            break;
    }

    if (shareUrl) {
        window.open(shareUrl, '_blank', 'width=600,height=400');
    }
}

function copyLink() {
    const url = window.location.href;

    navigator.clipboard.writeText(url).then(() => {
        // Show feedback
        const btn = event.currentTarget;
        const originalHTML = btn.innerHTML;

        btn.innerHTML = `
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="20 6 9 17 4 12"/>
            </svg>
        `;

        setTimeout(() => {
            btn.innerHTML = originalHTML;
        }, 2000);
    }).catch(err => {
        console.error('Failed to copy:', err);
        alert('Failed to copy link');
    });
}

// Code block enhancements
function enhanceCodeBlocks() {
    const codeBlocks = document.querySelectorAll('.doc-content pre code');

    codeBlocks.forEach(block => {
        const pre = block.parentElement;

        // Add copy button
        const copyBtn = document.createElement('button');
        copyBtn.className = 'code-copy-btn';
        copyBtn.innerHTML = 'Copy';
        copyBtn.onclick = () => {
            navigator.clipboard.writeText(block.textContent);
            copyBtn.innerHTML = 'Copied!';
            setTimeout(() => {
                copyBtn.innerHTML = 'Copy';
            }, 2000);
        };

        pre.style.position = 'relative';
        copyBtn.style.cssText = `
            position: absolute;
            top: 10px;
            right: 10px;
            padding: 6px 12px;
            background: var(--accent);
            color: white;
            border: none;
            border-radius: 4px;
            font-size: 0.8rem;
            cursor: pointer;
            opacity: 0.8;
            transition: opacity 0.2s;
        `;

        copyBtn.onmouseover = () => copyBtn.style.opacity = '1';
        copyBtn.onmouseout = () => copyBtn.style.opacity = '0.8';

        pre.appendChild(copyBtn);
    });
}

// Initialize on page load
document.addEventListener('DOMContentLoaded', () => {
    generateTableOfContents();
    initScrollSpy();
    enhanceCodeBlocks();

    // Handle initial hash
    if (window.location.hash) {
        setTimeout(() => {
            const target = document.querySelector(window.location.hash);
            if (target) {
                target.scrollIntoView({ behavior: 'smooth' });
            }
        }, 100);
    }
});

// Reading progress indicator (optional)
function initReadingProgress() {
    const progressBar = document.createElement('div');
    progressBar.style.cssText = `
        position: fixed;
        top: 0;
        left: 0;
        width: 0%;
        height: 3px;
        background: var(--accent);
        z-index: 9999;
        transition: width 0.1s ease;
    `;
    document.body.appendChild(progressBar);

    window.addEventListener('scroll', () => {
        const windowHeight = window.innerHeight;
        const documentHeight = document.documentElement.scrollHeight - windowHeight;
        const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
        const progress = (scrollTop / documentHeight) * 100;

        progressBar.style.width = `${progress}%`;
    });
}

// Uncomment to enable reading progress
// initReadingProgress();
