# transfertime
**Encoded Data Transfers for Reducing Raw Network Transfer Time**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of Technology and Applied Science (IJTAS)
- **Publication Date:** Oct, 2024
- **E-ISSN:**  2230-9004

### Abstract
Distributed systems frequently experience high network delays because large datasets are transmitted in raw form across network links. This study examines how encoded data transfers influence network transfer time and communication efficiency. By reducing the volume of transmitted data, encoded transfers shorten transmission duration and lower bandwidth consumption. Experimental evaluation across increasing file sizes shows that optimized and compressed transfers significantly reduce transfer time and improve scalability in distributed environments.

### Fundamental Contributions of the Work
- **Encoded Data Transfer Strategy:**  
Introduced a data transfer approach that converts raw payloads into compact encoded representations before transmission, reducing the amount of data sent across network links.

- **Efficient Network Communication Model:**  
Designed a data movement framework that reduces prolonged network occupancy by lowering payload size and shortening the duration of data transmission.

- **Distributed Transfer Simulation Framework:** 
Implemented a data transfer evaluation environment using Go based processing to simulate data movement and measure transfer time across varying file sizes.

- **Performance Evaluation with Increasing Data Volumes:**  
Analyzed transfer time behavior for file sizes ranging from 100 MB to 900 MB to study how optimized data movement improves scalability and communication efficiency.

### System Importance and Practical Impact
- **Reduced Network Transfer Time:**
Encoded data movement significantly decreases transmission duration compared with baseline raw data transfer mechanisms.

- **Improved Network Resource Utilization:**  
Lower data volume reduces bandwidth consumption and minimizes congestion on shared communication links.

- **Better Scalability for Data Intensive Systems:**  
Transfer time grows more gradually as data size increases, demonstrating improved scalability for large data transfers.

- **Applicability to Distributed Platforms:**  
The approach can benefit cloud systems, distributed storage platforms, analytics pipelines, and other environments that frequently exchange large datasets across networks.

### Experimental Results (Summary)

  | File Size (MB) | Unoptimized transfer time (ms) | Optimized transfer time (ms)| Improvement (%) |
  |----------------|--------------------------------| ----------------------------| ---------------|
  | 100            |  820                           | 480                         | 41.46          |
  | 300            |  2350                          | 1320                        | 43.83          |
  | 500            |  3800                          | 2050                        | 46.05          |
  | 700            |  5200                          | 2750                        | 47.12          |
  | 900            |  6650                          | 3400                        | 48.87          |

### Citation
Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing
* Naveen Kumar Bandaru
* International Journal of Technology and Applied Science (IJTAS) 
* E-ISSN 2230-9004
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijtas.com  \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







