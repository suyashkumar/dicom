# Grape Crush Data Visualization (2024 Errata)

This directory contains visualizations of the California Grape Crush data (2024 Errata).

## Data Source
The data was downloaded from [USDA NASS](https://www.nass.usda.gov/Statistics_by_State/California/Publications/Specialty_and_Other_Releases/Grapes/Crush/Errata/2024/2024_errata_gc_tb08.csv).

## Visualizations

### 1. Total Tons by Grape Type (`tons_by_grape_type.png`)
Shows the total tonnage crushed broken down by the main grape types: Wine Grapes (Red & White), Table Grapes, and Raisin Grapes.

### 2. Top 10 Varieties (`top_10_varieties.png`)
Displays the 10 specific grape varieties with the highest crushed tonnage.

### 3. Price vs. Quantity (`price_vs_tons.png`)
A scatter plot comparing the total tons crushed (on a log scale) against the average price per ton. This helps visualize the relationship between scarcity and value.

### 4. Production by District (`tons_by_district.png`)
Shows the total tonnage crushed in each of California's grape pricing districts.

## How to Reproduce
Run the python script `analyze_grapes.py` to regenerate the images.
```bash
pip install pandas matplotlib seaborn
python3 analyze_grapes.py
```
