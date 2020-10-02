/**
    CSE-310 P1: Take input and output encoded version of the input
    @file encode.cpp
    @author Austin Spencer
    @version Milestone 09/15/20
*/
#include <iostream>
#include <string>

using namespace std;

/**
    Sorts string array via insertion sort
    @param str2 pointer to the string array
    @param n number of strings in array
*/
void insertionSort(string *str2, int n)
{
    int x, y;
    string key;
    for (x = 1; x < n; x++)
    {
        key = str2[x];
        y = x - 1;
        while (y >= 0 && str2[y] > key)
        {
            str2[y + 1] = str2[y];
            y = y - 1;
        }
        str2[y + 1] = key;
    }
}

/**
    Swap helper function
    @param a pointer to the smaller string
    @param b pointer to the larger string
*/
void swap(string *a, string *b)
{
    string t = *a;
    *a = *b;
    *b = t;
}

/**
    Partition using last string as pivot
    @param str2 string array
    @param left smaller integer
    @param right larger integer
*/
int partition(string str2[], int left, int right)
{
    string pivot = str2[right]; // pivot
    int i = (left - 1);

    for (int j = left; j <= right - 1; j++)
    {
        // if current string is smaller than pivot, increment the left string
        if (str2[j] <= pivot)
        {
            i++; // increment index of smaller string
                 // swap strings at i and j
            swap(&str2[i], &str2[j]);
        }
    }
    // swap in pivot and return its position
    swap(&str2[i + 1], &str2[right]);
    return (i + 1);
}

/**
    Sorts given string array via quicksort
    @param str2 string array
    @param left starting index
    @param right largest index
*/
void quickSort(string str2[], int left, int right)
{
    if (left < right)
    {
        // partition the str2
        int pivot = partition(str2, left, right);

        // sort the sub str2 (divide & conquer)
        quickSort(str2, left, pivot - 1);
        quickSort(str2, pivot + 1, right);
    }
}

/**
    Formats our output and prints the encoded line with std::cout
    @param last pointer to string of each final character of sorted string array
    @param n number of characters in string last
*/
int printEncodedLine(string *last, int n, int clusters)
{
    int num = 1;
    // iterate through entire string array
    for (int i = 0; i < n; i++)
    {
        // if next letter is the same as current just increment num and continue for loop
        if (last[i] == last[i + 1])
        {
            num++;
            continue;
        }
        clusters++;
        cout << num << " " << last[i];

        // if statement to avoid printing extra space at the end
        if (i != (n - 1))
        {
            cout << " ";
        }

        // reset num to 1
        num = 1;
    }
    return clusters;
}

void calculateCompression(int clusters, int t) {
    cout << endl << "t: " << t << " clust: " << clusters << endl;
    int compRatio = ((float)(t - clusters)/(float)t)*100;
    cout << "Compression Ratio: " << compRatio << "%" << endl;
}

    /**
    The main function; reads input and calls other functions
    @param argc number of command line arguments
    @param argv array of each argument
*/
    int
    main(int argc, char **argv)
{
    // Set variable keyword as the value of our second cmd line argument
    string keyword = argv[1];
    string str;
    int clusters = 0;
    int t = 0;
    int lines = 0;

    // While there is input in the txt file, set str equal to the current line
    while (getline(cin, str))
    {
        // If the line isn't empty, encode the line
        if (!str.empty())
        {
            if (lines != 0)
            {
                cout << endl;
            }
            else
            {
                // Increment number of lines, this will allow to skip first line and add EOL to the rest
                lines++;
            }

            int n = str.size();
            t += n;
            // Initialize str2 (string array we will be sorting)
            string str2[n];

            string original = str;
            str2[0] = str;

            for (string::size_type i = 1; i < (n); i++)
            {
                /**
                shift given string:
                Take substring(not including first char)
                Take substring of only first char and add it to the end of first substring
                */
                str = str.substr(1, (n - 1)) + str.substr(0, 1);
                // append each shift into str2 array
                str2[i] = str;
            }

            // Check keyword for which sorting algorithm to use
            if (keyword == "quick")
            {
                quickSort(str2, 0, n - 1);
            }
            else if (keyword == "insertion")
            {
                insertionSort(&str2[0], n); // Pass the address of our string array
            }
            else
            {
                cout << "invalid command line argument" << endl;
                return -1; // return non-zero since an error occurred
            }

            int originalLocation;
            string current;
            string last[n];

            for (int i = 0; i < n; ++i)
            {
                current = str2[i];
                // append our string last with the final character in str2
                last[i] = current[(n - 1)];
                // check if the current str2 is our original string
                if (str2[i] == original)
                {
                    // If it is our original string set originalLocation equal to the index we are at
                    originalLocation = i;
                }
            }
            // Print the index of our original string
            cout << originalLocation << endl;
            clusters = printEncodedLine(last, n, clusters);
        }
        else
        {
            // If the line is empty, print an empty line
            cout << endl;
        }
    }

    if (cin.bad())
    {
        cout << "IO error" << endl;
    }
    else if (!cin.eof())
    {
        cout << "format error" << endl;
    }
    else
    {
        if (str.empty())
        {
            // Add empty line if file ends with an empty line
            cout << endl;
        }
        // end of file
    }
    cout << endl;
    calculateCompression(clusters, t);

    return 0;
} // End of program