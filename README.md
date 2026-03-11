# transfertime
**Encoded Data Transfers for Reducing Raw Network Transfer Time**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of Technology and Applied Science (IJTAS)
- **Publication Date:** Oct, 2024
- **ISSN:**  2230-9004

### Abstract
Distributed transaction systems often experience high commit latency because each transaction is processed independently with repeated coordination and synchronization among nodes. This work examines the impact of immediate commit processing on latency as cluster size increases. A batching based commit approach is introduced where multiple transactions are processed collectively to reduce repeated coordination overhead. Experimental analysis across different cluster sizes shows that grouped commit processing significantly lowers commit latency and improves scalability in distributed environments.

### Primary Contributions of the Study
- **Collective Transaction Commit Strategy:**  
Presented a commit handling strategy that groups several transactions together, allowing multiple operations to complete within a single coordination cycle and reducing repeated commit procedures.

- **Efficient Coordination Handling:**  
Developed a commit processing model that reduces repeated communication and synchronization between coordinator and participant nodes during distributed transaction finalization.
- **Prototype Based Experimental Setup:** 
Created a distributed transaction processing simulation using Go based concurrent processes to emulate coordinator participant interactions and observe commit behavior under varying workloads.

- **Evaluation with Increasing Node Configurations:**  
Analyzed commit latency across clusters containing 3, 5, 7, 9, and 11 nodes to understand how transaction completion time changes as distributed systems expand.

### System Relevance and Operational Value
- **Faster Transaction Completion:**
Grouped commit processing shortens the time required to finalize transactions by reducing the number of coordination rounds compared with traditional independent commit mechanisms.

- **Improved Processing Efficiency:**  
Handling several transactions together decreases repeated synchronization and communication overhead, enabling faster processing of transactional workloads in distributed environments.

- **Stable Performance with System Growth:**  
Latency increases gradually as cluster size expands, demonstrating that batching commit operations supports scalable transaction processing in distributed infrastructures.

- **Usefulness for Modern Distributed Platforms:**  
The approach can support distributed databases, cloud services, financial applications, and large scale enterprise platforms requiring efficient and scalable transaction commit processing.

### Experimental Results (Summary)

  | Nodes | Lock Based 2PL CPU %| Lightweight Runtime Detection %| Improvment (%) |
  |-------|---------------------| -------------------------------| ---------------|
  | 3     |  88                 | 55                             | 37.50          |
  | 5     |  84                 | 49                             | 41.67          |
  | 7     |  82                 | 46                             | 43.90          |
  | 9     |  80                 | 43                             | 46.25          |
  | 11    |  79                 | 41                             | 48.10          |

### Citation
Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing
* Naveen Kumar Bandaru
* International Journal of Intelligent Systems and Applications in Engineering (IJISAE) 
* ISSN 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijisae.org/index.php/IJISAE \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







