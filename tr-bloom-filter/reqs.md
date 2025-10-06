**High-Level Requirements: Bloom Filter Implementation**

**Objective:**
Design and implement a minimal, working Bloom Filter from scratch (no libraries) to understand probabilistic data structures and hashing-based membership tests.

---

### 1. Functional Requirements

* Must allow insertion of elements into the filter.
* Must allow membership checks (“possibly in set” / “definitely not in set”).
* Must not support deletion (unless implementing a Counting Bloom variant, which is out of scope for first version).
* Must be space-efficient and faster than storing all elements explicitly.

---

### 2. Core Components

* **Bit Array:** Fixed-size array to represent filter state.
* **Hash Functions:** Multiple deterministic hash functions mapping input to bit positions.
* **Bit Setting Logic:** Set bits on insertion, check all bits on lookup.
* **False Positives:** Acknowledge their presence; ensure no false negatives.

---

### 3. Configuration Inputs

* Expected number of elements (n)
* Desired false positive rate (p)
* Derived parameters:

  * Bit array size (m)
  * Number of hash functions (k)

Formulas to derive m and k are standard Bloom Filter equations.

---

### 4. Operations

**Initialize:** Create zeroed bit array, compute optimal parameters.
**Insert(item):** Apply k hash functions → set bits.
**Contains(item):** Check same k bit positions → decide membership.

---

### 5. Performance Requirements

* O(k) time per insert and query.
* O(m) space complexity.
* Deterministic behavior for identical input sequence.

---

### 6. Error Conditions

* Must handle exceeding expected n gracefully (with higher false positive rate).
* Must handle non-hashable inputs (e.g., complex objects) via explicit error or serialization strategy.

---

### 7. Testing Goals

* Verify correctness for known included/excluded items.
* Measure false positive rate empirically against theoretical target.
* Evaluate scaling behavior for different n, p combinations.

---

### 8. Extensions (optional after base version)

* Counting Bloom Filter (support deletes).
* Scalable Bloom Filter (dynamic resizing).
* Persistent Bloom Filter (disk-backed or serialized).
* Cryptographic hash variants for security-focused use.

---

### 9. Reference Materials


1. **Original paper:** Burton H. Bloom, *“Space/Time Trade-offs in Hash Coding with Allowable Errors”* (1970).
1. **Article:** https://systemdesign.one/bloom-filters-explained/ 
2. **MIT 6.006 Lecture Notes:** Probabilistic data structures.
3. **Stanford CS166:** *Bloom Filters and Cuckoo Filters* lecture.
4. **“Probabilistic Data Structures and Algorithms for Big Data Applications”* by Andrei Broder & Michael Mitzenmacher.
5. **Wikipedia:** *Bloom filter* page for mathematical background and parameter derivations.
