import java.io.*;
import java.util.*;

public class Program {
    static class Pair {
        int first;
        int second;

        public Pair(int first, int second) {
            this.first = first;
            this.second = second;
        }
    }

    static void build(List<List<Pair>> graph, int[] depth, List<List<Pair>> dp, int max_i, int v, int p, int c) {
        if (p != v) depth[v] = depth[p] + 1;

        dp.get(v).add(0, new Pair(p, c));
        for (int i = 1; i <= max_i; ++i) {
            dp.get(v).add(i, new Pair(dp.get(dp.get(v).get(i - 1).first).get(i - 1).first, Math.min(dp.get(v).get(i - 1).second, dp.get(dp.get(v).get(i - 1).first).get(i - 1).second)));
        }

        for (Pair next : graph.get(v)) build(graph, depth, dp, max_i, next.first, v, next.second);
    }

    static int min_lca(int[] depth, List<List<Pair>> dp, int max_i, int u, int v) {
        int result = Integer.MAX_VALUE;

        if (depth[v] > depth[u]) {
            int temp = u;
            u= v;
            v = temp;
        }

        for (int i = max_i; i >= 0; --i) {
            if (depth[dp.get(u).get(i).first] >= depth[v]) {
                result = Math.min(result, dp.get(u).get(i).second);
                u = dp.get(u).get(i).first;
            }
        }

        if (v == u) return result;

        for (int i = max_i; i >= 0; --i) {
            if (dp.get(v).get(i).first!= dp.get(u).get(i).first) {
                result = Math.min(result, dp.get(v).get(i).second);
                v = dp.get(v).get(i).first;
                result = Math.min(result, dp.get(u).get(i).second);
                u = dp.get(u).get(i).first;
            }
        }

        result = Math.min(result, dp.get(u).get(0).second);
        result = Math.min(result, dp.get(v).get(0).second);

        return result;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        int n = Integer.parseInt(br.readLine());
        List<List<Pair>> graph = new ArrayList<>();
        for (int i = 0; i < n; ++i) graph.add(new ArrayList<>());

        for (int i = 1; i < n; ++i) {
            String[] line = br.readLine().split(" ");
            int x = Integer.parseInt(line[0]);
            int y = Integer.parseInt(line[1]);
            graph.get(x).add(new Pair(i, y));
        }

        int[] depth = new int[n];
        List<List<Pair>> dp = new ArrayList<>();
        for (int i = 0; i < n; ++i) dp.add(new ArrayList<>());
        int max_i = 1;
        while ((1 << max_i) <= n) ++max_i;
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j <= max_i; ++j) dp.get(i).add(new Pair(0, 0));
        }

        build(graph, depth, dp, max_i, 0, 0, Integer.MAX_VALUE);

        int m = Integer.parseInt(br.readLine());
        for (int req = 0; req < m; ++req) {
            String[] line = br.readLine().split(" ");
            int u = Integer.parseInt(line[0]);
            int v = Integer.parseInt(line[1]);
            bw.write(String.valueOf(min_lca(depth, dp, max_i, u, v)) + "\n");
        }

        bw.flush();
    }
}