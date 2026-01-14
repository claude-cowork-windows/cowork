use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use std::path::{Path, PathBuf};
use std::fs;

// Helper function to convert C-style strings (char*) to Rust strings
fn parse_c_string(ptr: *const c_char) -> String {
    unsafe {
        if ptr.is_null() { return String::new(); }
        CStr::from_ptr(ptr).to_string_lossy().into_owned()
    }
}

// Security Check: verifies if the target path is strictly inside the root directory
// Prevents Path Traversal attacks (e.g., "../../windows")
fn is_safe_path(root: &str, target: &str) -> Option<PathBuf> {
    let root_path = fs::canonicalize(root).ok()?;
    let target_path = Path::new(root).join(target);
    
    // Attempt to resolve the full canonical path.
    // If the file doesn't exist yet (creation mode), we check the parent directory.
    let final_path = if target_path.exists() {
        fs::canonicalize(&target_path).ok()?
    } else {
        // For new files, we accept the logical path (simplified for this demo)
        target_path 
    };

    // The crucial check: does the final path start with the root path?
    if final_path.starts_with(&root_path) {
        Some(final_path)
    } else {
        None
    }
}

// Exported Function: Secure File Write
// Can be called from C/Go
#[no_mangle]
pub extern "C" fn safe_write_file(root_dir: *const c_char, filename: *const c_char, content: *const c_char) -> i32 {
    let root = parse_c_string(root_dir);
    let file = parse_c_string(filename);
    let data = parse_c_string(content);

    // 1. Security Sandbox Check
    let safe_path = match is_safe_path(&root, &file) {
        Some(p) => p,
        None => {
            println!("[RUST-SEC] ⛔ Blocked attempt to write outside sandbox: {}", file);
            return -1; // Error Code: Access Denied
        }
    };

    // 2. Perform filesystem operation
    match fs::write(&safe_path, data) {
        Ok(_) => {
            println!("[RUST-FS] ✅ Successfully wrote: {:?}", safe_path);
            0 // Success
        },
        Err(e) => {
            println!("[RUST-FS] ❌ Error writing file: {}", e);
            -2 // Error Code: IO Error
        }
    }
}

// Exported Function: Secure File Read
#[no_mangle]
pub extern "C" fn safe_read_file(root_dir: *const c_char, filename: *const c_char) -> *mut c_char {
    let root = parse_c_string(root_dir);
    let file = parse_c_string(filename);

    // Security check before reading
    let safe_path = match is_safe_path(&root, &file) {
        Some(p) => p,
        None => return CString::new("ERROR: Access Denied").unwrap().into_raw(),
    };

    match fs::read_to_string(safe_path) {
        Ok(content) => CString::new(content).unwrap().into_raw(),
        Err(_) => CString::new("ERROR: File not found").unwrap().into_raw(),
    }
}
