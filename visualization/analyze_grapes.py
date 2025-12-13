import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import os

def load_and_clean_data(filepath):
    """Loads the CSV and filters for summary rows."""
    try:
        df = pd.read_csv(filepath)
        # Filter for summary rows to get aggregated data
        summary_df = df[df['row_type_name'] == 'summary_row'].copy()

        # Ensure numeric columns are actually numeric
        summary_df['tons'] = pd.to_numeric(summary_df['tons'], errors='coerce')
        summary_df['wt_price'] = pd.to_numeric(summary_df['wt_price'], errors='coerce')

        return summary_df
    except Exception as e:
        print(f"Error loading data: {e}")
        return None

def plot_tons_by_grape_type(df, output_dir):
    """Bar chart of total tons by grape type."""
    plt.figure(figsize=(10, 6))

    # Group by grape type
    type_data = df.groupby('grape_type_name')['tons'].sum().sort_values(ascending=False).reset_index()

    sns.barplot(data=type_data, x='grape_type_name', y='tons', hue='grape_type_name', palette='viridis', legend=False)
    plt.title('Total Tons Crushed by Grape Type (2024)')
    plt.xlabel('Grape Type')
    plt.ylabel('Total Tons')
    plt.xticks(rotation=45)
    plt.tight_layout()

    output_path = os.path.join(output_dir, 'tons_by_grape_type.png')
    plt.savefig(output_path)
    print(f"Saved {output_path}")
    plt.close()

def plot_top_varieties(df, output_dir):
    """Horizontal bar chart of top 10 varieties by tons."""
    plt.figure(figsize=(12, 8))

    # Group by variety
    variety_data = df.groupby('variety_name')['tons'].sum().sort_values(ascending=False).head(10).reset_index()

    sns.barplot(data=variety_data, y='variety_name', x='tons', hue='variety_name', palette='magma', legend=False)
    plt.title('Top 10 Grape Varieties by Tons Crushed (2024)')
    plt.xlabel('Total Tons')
    plt.ylabel('Variety')
    plt.tight_layout()

    output_path = os.path.join(output_dir, 'top_10_varieties.png')
    plt.savefig(output_path)
    print(f"Saved {output_path}")
    plt.close()

def plot_price_vs_tons(df, output_dir):
    """Scatter plot of Average Price vs Total Tons."""
    plt.figure(figsize=(12, 8))

    # Group by variety to get overall totals and weighted average price
    # We need to recalculate weighted average price across districts

    # Calculate total value per row
    df['total_value'] = df['tons'] * df['wt_price']

    variety_stats = df.groupby('variety_name').agg({
        'tons': 'sum',
        'total_value': 'sum',
        'grape_type_name': 'first' # Assuming a variety belongs to one type
    }).reset_index()

    variety_stats['avg_price'] = variety_stats['total_value'] / variety_stats['tons']

    # Filter out very small productions to reduce noise? Maybe keep all.

    sns.scatterplot(data=variety_stats, x='tons', y='avg_price', hue='grape_type_name', alpha=0.7)
    plt.xscale('log') # Log scale for tons as it likely spans orders of magnitude
    plt.title('Average Price vs. Total Tons (by Variety)')
    plt.xlabel('Total Tons (Log Scale)')
    plt.ylabel('Average Price ($/Ton)')
    plt.legend(title='Grape Type')
    plt.grid(True, which="both", ls="-", alpha=0.2)
    plt.tight_layout()

    output_path = os.path.join(output_dir, 'price_vs_tons.png')
    plt.savefig(output_path)
    print(f"Saved {output_path}")
    plt.close()

def plot_district_production(df, output_dir):
    """Bar chart of total tons by district."""
    plt.figure(figsize=(14, 6))

    district_data = df.groupby('district')['tons'].sum().reset_index()
    district_data['district'] = district_data['district'].astype(str) # Ensure categorical plotting

    # Sort by district number (convert to int for sorting)
    district_data['district_int'] = district_data['district'].astype(int)
    district_data = district_data.sort_values('district_int')

    sns.barplot(data=district_data, x='district', y='tons', color='skyblue')
    plt.title('Total Tons Crushed by District (2024)')
    plt.xlabel('District')
    plt.ylabel('Total Tons')
    plt.tight_layout()

    output_path = os.path.join(output_dir, 'tons_by_district.png')
    plt.savefig(output_path)
    print(f"Saved {output_path}")
    plt.close()

def main():
    # Use paths relative to the script location
    script_dir = os.path.dirname(os.path.abspath(__file__))
    filepath = os.path.join(script_dir, 'grape_crush.csv')
    output_dir = script_dir

    if not os.path.exists(filepath):
        print(f"File not found: {filepath}")
        return

    df = load_and_clean_data(filepath)
    if df is not None:
        plot_tons_by_grape_type(df, output_dir)
        plot_top_varieties(df, output_dir)
        plot_price_vs_tons(df, output_dir)
        plot_district_production(df, output_dir)
        print("All visualizations created successfully.")

if __name__ == "__main__":
    main()
