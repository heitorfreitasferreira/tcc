// Dropdown toggle for top navigation bar.
// Closes any open dropdown when clicking outside.

let openDropdown = null;

function toggleDropdown(id) {
  const dd = document.getElementById(id);
  if (!dd) return;

  const isOpen = dd.classList.contains("open");
  closeAll();

  if (!isOpen) {
    dd.classList.add("open");
    openDropdown = dd;
  }
}

function closeAll() {
  if (openDropdown) {
    openDropdown.classList.remove("open");
    openDropdown = null;
  }
}

document.addEventListener("click", (e) => {
  if (openDropdown && !openDropdown.contains(e.target)) {
    closeAll();
  }
});

// Close dropdowns after HTMX swaps
document.addEventListener("htmx:afterSettle", () => {
  closeAll();
});

// Expose globally for onclick handlers
window.toggleDropdown = toggleDropdown;
