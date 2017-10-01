package Recongize


public class PersonRecognition {
	public static boolean Recognition(List<Vertex> pWordSegResult, WordNet wordNetOptimum, WordNet wordNetAll)
	{
	List<EnumItem<NR>> roleTagList = roleObserve(pWordSegResult);
	if (HanLP.Config.DEBUG)
	{
	StringBuilder sbLog = new StringBuilder();
	Iterator<Vertex> iterator = pWordSegResult.iterator();
	for (EnumItem<NR> nrEnumItem : roleTagList)
	{
	sbLog.append('[');
	sbLog.append(iterator.next().realWord);
	sbLog.append(' ');
	sbLog.append(nrEnumItem);
	sbLog.append(']');
	}
	System.out.printf("人名角色观察：%s\n", sbLog.toString());
	}
	List<NR> nrList = viterbiComputeSimply(roleTagList);
	if (HanLP.Config.DEBUG)
	{
	StringBuilder sbLog = new StringBuilder();
	Iterator<Vertex> iterator = pWordSegResult.iterator();
	sbLog.append('[');
	for (NR nr : nrList)
	{
	sbLog.append(iterator.next().realWord);
	sbLog.append('/');
	sbLog.append(nr);
	sbLog.append(" ,");
	}
	if (sbLog.length() > 1) sbLog.delete(sbLog.length() - 2, sbLog.length());
	sbLog.append(']');
	System.out.printf("人名角色标注：%s\n", sbLog.toString());
	}

	PersonDictionary.parsePattern(nrList, pWordSegResult, wordNetOptimum, wordNetAll);
	return true;
}

/**
 * 角色观察(从模型中加载所有词语对应的所有角色,允许进行一些规则补充)
 * @param wordSegResult 粗分结果
 * @return
 */
public static List<EnumItem<NR>> roleObserve(List<Vertex> wordSegResult){
	List<EnumItem<NR>> tagList = new LinkedList<EnumItem<NR>>();
	Iterator<Vertex> iterator = wordSegResult.iterator();
	iterator.next();
	tagList.add(new EnumItem<NR>(NR.A, NR.K)); //  始##始 A K
	while (iterator.hasNext())
	{
	Vertex vertex = iterator.next();
	EnumItem<NR> nrEnumItem = PersonDictionary.dictionary.get(vertex.realWord);
	if (nrEnumItem == null){
		switch (vertex.guessNature())
		{
		case nr:
		{
		// 有些双名实际上可以构成更长的三名
		if (vertex.getAttribute().totalFrequency <= 1000 && vertex.realWord.length() == 2)
		{
		nrEnumItem = new EnumItem<NR>(NR.X, NR.G);
		}
		else nrEnumItem = new EnumItem<NR>(NR.A, PersonDictionary.transformMatrixDictionary.getTotalFrequency(NR.A));
		}break;
		case nnt:
		{
		// 姓+职位
		nrEnumItem = new EnumItem<NR>(NR.G, NR.K);
		}break;
		default:
		{
		nrEnumItem = new EnumItem<NR>(NR.A, PersonDictionary.transformMatrixDictionary.getTotalFrequency(NR.A));
		}break;
		}
		}
	tagList.add(nrEnumItem);
	}
	return tagList;
}

/**
 * 维特比算法求解最优标签
 * @param roleTagList
 * @return
 */
public static List<NR> viterbiCompute(List<EnumItem<NR>> roleTagList){
	return Viterbi.computeEnum(roleTagList, PersonDictionary.transformMatrixDictionary);
}

/**
 * 简化的"维特比算法"求解最优标签
 * @param roleTagList
 * @return
 */
public static List<NR> viterbiComputeSimply(List<EnumItem<NR>> roleTagList){
	return Viterbi.computeEnumSimply(roleTagList, PersonDictionary.transformMatrixDictionary);
}
}


/**
 * 音译人名识别
 */

/**
 * 执行识别
 * @param segResult 粗分结果
 * @param wordNetOptimum 粗分结果对应的词图
 * @param wordNetAll 全词图
 */
public static void Recognition(List<Vertex> segResult, WordNet wordNetOptimum, WordNet wordNetAll) {
	StringBuilder sbName = new StringBuilder();
	int appendTimes = 0;
	ListIterator<Vertex> listIterator = segResult.listIterator();
	listIterator.next();
	int line = 1;
	int activeLine = 1;
	while (listIterator.hasNext())
	{
	Vertex vertex = listIterator.next();
	if (appendTimes > 0) {
		if (vertex.guessNature() == Nature.nrf || TranslatedPersonDictionary.containsKey(vertex.realWord))
		{
			sbName.append(vertex.realWord);
			++appendTimes;
		}
		else
		{
			// 识别结束
			if (appendTimes > 1)
			{
			if (HanLP.Config.DEBUG)
			{
			System.out.println("音译人名识别出：" + sbName.toString());
			}
			wordNetOptimum.insert(activeLine, new Vertex(Predefine.TAG_PEOPLE, sbName.toString(), new CoreDictionary.Attribute(Nature.nrf), WORD_ID), wordNetAll);
			}
			sbName.setLength(0);
			appendTimes = 0;
		}
	}
	else
	{
		// nrf触发识别
		if (vertex.guessNature() == Nature.nrf
		//                        || TranslatedPersonDictionary.containsKey(vertex.realWord)
		)
		{
			sbName.append(vertex.realWord);
			++appendTimes;
			activeLine = line;
		}
	}

	line += vertex.realWord.length();
	}
}
